FROM golang:1.13
WORKDIR /go/src/app
COPY . .

RUN go build -mod=vendor main.go

ENTRYPOINT ["main"]
