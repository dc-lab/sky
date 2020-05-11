FROM    golang:1.14.2-alpine3.11 AS builder
RUN     apk update && apk add bash ca-certificates git gcc g++ libc-dev binutils file
WORKDIR /usr/src/sky/agent/
COPY    src ./
COPY    docker/config.json /config.json
COPY    docker/run.sh /run.sh
RUN     go mod download
RUN     go build -o /agent

FROM alpine:3.11 AS production
RUN  apk update && apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*
WORKDIR /usr/lib/sky/agent/lib/
COPY --from=builder /agent ./
COPY --from=builder /config.json ./
COPY --from=builder /run.sh ./
CMD  ["./run.sh"]
