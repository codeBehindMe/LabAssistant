FROM golang:1.14

WORKDIR /

RUN mkdir assistant

WORKDIR assistant

COPY . .

ENV GO111MODULES=on

RUN go build -o assistant main.go

ENTRYPOINT ["./assistant"]
