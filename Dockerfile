FROM    golang:1.15.2-alpine3.12 AS builder
RUN     apk update && apk add ca-certificates git libc-dev
WORKDIR /usr/src/app/sky/
COPY    go.mod .
COPY    go.sum .
RUN     go mod download
COPY    . .
RUN     go build -o / ./...


FROM alpine:3.12 AS production
RUN  apk update && apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*


FROM production as binary
ARG executable
COPY --from=builder /${executable} ./app
CMD  ["./app"]
