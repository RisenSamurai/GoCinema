FROM golang:1.22-alpine

WORKDIR /app

COPY rating-service/go.mod rating-service/go.sum ./

RUN go mod download

COPY rating-service .

RUN go build -o main .

EXPOSE 8081

CMD ["./main"]