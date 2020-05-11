FROM    golang:latest
WORKDIR /Users/ilpauzner/go/src
#/dc-store/auth
COPY    . .
RUN go build main.go
RUN chmod 755 wait-for-it.sh
CMD  ["./wait-for-it.sh", "--timeout=60", "https://yandex.ru/", "--", "./main"]
# example of waiting to, maybe useful later.