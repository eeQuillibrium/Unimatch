Backend for matching application



==DEV==
protoc --go_out=../gen/go --go_opt=paths=source_relative --go-grpc_out=../gen/go --go-grpc_opt=paths=source_relative protobuf_services/auth/auth.proto
migrate -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable -path migrations up