FROM golang:latest

COPY ./ ./

RUN go mod download

RUN go build -o ./auth_service/bin ./auth_service/cmd/main.go