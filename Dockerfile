FROM golang:1.23-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

CMD ["./main"]
