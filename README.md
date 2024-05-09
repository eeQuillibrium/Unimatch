Backend for matching application
Now it works with ngrok
.==================================================================.
protoc --go_out=../gen/go --go_opt=paths=source_relative --go-grpc_out=../gen/go --go-grpc_opt=paths=source_relative protobuf_services/auth/auth.proto
==========================DOCKER=====================
1. git clone {repolink}
2. cd Unimatch
3. docker-compose up
4. open another terminal, cd to ./auth_service
5. migrate -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable -path migrations up