FROM golang:1.22-alpine AS builder
WORKDIR /build
COPY go.mod main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o relay .

FROM scratch
COPY --from=builder /build/relay /relay
ENTRYPOINT ["/relay"]
