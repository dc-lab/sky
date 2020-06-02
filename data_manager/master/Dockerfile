FROM    golang:1.14.2-alpine3.11 AS builder
RUN     apk update && apk add bash ca-certificates git gcc g++ libc-dev binutils file
WORKDIR /usr/src/app/sky/data_manager/
COPY    . .
RUN     go mod download
RUN     go build -o /app .
RUN     cp -r migrations /migrations

FROM alpine:3.11 AS production
RUN  apk update && apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*
COPY --from=builder /app ./
COPY --from=builder /migrations ./migrations
CMD  ["./app"]
