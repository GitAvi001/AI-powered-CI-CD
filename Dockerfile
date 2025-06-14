FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o AI-powered-CI-CD ./cmd/api

EXPOSE 8080

CMD ["./AI-powered-CI-CD"]