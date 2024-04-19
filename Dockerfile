FROM golang:latest

COPY ./ ./
RUN go mod download

RUN go build -o ./api_gateway_service/bin ./api_gateway_service/cmd/main.go
RUN go build -o ./auth_service/bin ./auth_service/cmd/main.go

CMD ["migrate -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable -path auth_service/migrations up"]