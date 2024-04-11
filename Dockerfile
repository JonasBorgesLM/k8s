FROM golang:1.20.3-alpine
WORKDIR /app

RUN go mod init server
COPY server.go .

RUN go build -o server
CMD [ "./server" ]