package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	listenPort := getEnv("LISTEN_PORT", "8080")
	targetHost := mustEnv("TARGET_HOST")
	targetPort := getEnv("TARGET_PORT", listenPort)

	target := net.JoinHostPort(targetHost, targetPort)

	ln, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		log.Fatalf("listen :%s: %v", listenPort, err)
	}
	log.Printf("relaying :%s -> %s", listenPort, target)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept: %v", err)
			continue
		}
		go relay(conn, target)
	}
}

func relay(src net.Conn, target string) {
	defer src.Close()

	dst, err := net.Dial("tcp", target)
	if err != nil {
		log.Printf("dial %s: %v", target, err)
		return
	}
	defer dst.Close()

	done := make(chan struct{}, 2)

	pipe := func(w, r net.Conn) {
		io.Copy(w, r)
		// half-close so the other side sees EOF
		if tc, ok := w.(*net.TCPConn); ok {
			tc.CloseWrite()
		}
		done <- struct{}{}
	}

	go pipe(dst, src)
	go pipe(src, dst)

	<-done
	<-done
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("required env var %s is not set", key)
	}
	return v
}
