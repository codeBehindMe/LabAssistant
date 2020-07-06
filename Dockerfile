FROM golang:1.14

WORKDIR /

RUN mkdir app

WORKDIR app

COPY . .

ENV GO111MODULES=on

RUN go build -o app main.go

ENTRYPOINT ["./app"]
