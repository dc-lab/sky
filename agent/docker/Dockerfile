FROM    golang:1.14.2-alpine3.11 AS builder
RUN     apk update && apk add ca-certificates git libc-dev
WORKDIR /usr/src/sky/agent/
COPY    src ./
RUN     go mod download
RUN     go build -o /sky-agent

FROM    alpine:3.11 AS production
RUN     apk update && apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*
COPY    docker/etc /etc
COPY    docker/usr /usr
COPY    --from=builder /sky-agent /usr/bin/sky-agent
COPY    --from=ochinchina/supervisord:latest /usr/local/bin/supervisord /usr/bin/supervisord
CMD     ["/usr/lib/sky/agent/run.sh"]

HEALTHCHECK --interval=5s --timeout=1s --start-period=10s --retries=1 \
    CMD /usr/lib/sky/agent/check-agent-health.sh || exit 1
