FROM golang:latest
WORKDIR /go/src/app
COPY main.go .
RUN go mod init example.com/m
RUN go build -o myapp .
CMD ["./myapp"]