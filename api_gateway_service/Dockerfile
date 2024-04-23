FROM golang:latest

COPY ./ ./

RUN go mod download

RUN go build -o ./api_gateway_service/bin ./api_gateway_service/cmd/main.go
