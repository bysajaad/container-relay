# container-relay

A minimal TCP relay proxy written in Go, deployable on Docker. Designed as a front-door relay for v2ray (or any TCP service) — listens on a port and transparently forwards all traffic to a configured upstream.

## Usage

### docker run

```sh
docker run -d \
  -e TARGET_HOST=your-v2ray-server.example.com \
  -e TARGET_PORT=443 \
  -e LISTEN_PORT=443 \
  -p 443:443 \
  ghcr.io/bysajaad/container-relay
```

### docker compose

Edit `docker-compose.yml` with your target, then:

```sh
docker compose up -d
```

## Environment variables

| Variable | Required | Default | Description |
|---|---|---|---|
| `TARGET_HOST` | yes | — | Upstream host to forward connections to |
| `TARGET_PORT` | no | same as `LISTEN_PORT` | Upstream port |
| `LISTEN_PORT` | no | `8080` | Port the relay listens on inside the container |

## Build

```sh
docker build -t container-relay .
```

The final image is built `FROM scratch` — only the statically compiled binary is included (~3 MB).
