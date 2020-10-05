FROM    golang:1.14.2-alpine3.11 AS builder
RUN     apk update && apk add ca-certificates git libc-dev
WORKDIR /usr/src/app/sky/node/
COPY    go.mod .
COPY    go.sum .
RUN     go mod download
COPY    . .
RUN     go build -o /app .

FROM alpine:3.11 AS production
RUN  apk update && apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*
COPY --from=builder /app ./
CMD  ["./app"]
