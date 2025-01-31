FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o server

CMD ["./server"]

EXPOSE 8080
