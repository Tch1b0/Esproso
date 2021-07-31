FROM golang:latest

WORKDIR /go/src/app

COPY ./main.go ./main.go
COPY ./go.mod ./go.mod
COPY ./data ./data
COPY ./brain ./brain

RUN go mod download
RUN go build github.com/Tch1b0/Esproso

CMD ["./Esproso"]