FROM golang:1.18.4-alpine AS builder
WORKDIR /go/src/httpx

ADD . .
RUN make build

FROM alpine:3.16.1
RUN apk -U upgrade --no-cache \
    && apk add --no-cache bind-tools ca-certificates
COPY --from=builder /go/src/httpx/httpx /usr/local/bin/

ENTRYPOINT ["httpx", "-silent", "-json", "-no-fallback", "-pipeline", "-response-in-json", "-tech-detect", "-output", "/output"]
