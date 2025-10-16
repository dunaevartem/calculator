FROM golang:1.24.9

WORKDIR /go/src/app

COPY . .

RUN go build -o app

EXPOSE 8080

CMD ["./app"]
