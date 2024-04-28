Backend for matching application

==DEV==
protoc --go_out=../gen/go --go_opt=paths=source_relative --go-grpc_out=../gen/go --go-grpc_opt=paths=source_relative protobuf_services/auth/auth.proto
==========================ANDREY=====================
1. git clone {repolink}
2. cd Unimatch
3. docker-compose up
4. open another terminal, cd to ./auth_service
5. migrate -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable -path migrations up  

localhost:8080/auth/signUp - регистрация (POST)  
localhost:8080/auth/signIn - авторизация (POST)  
Такой json:  
{  
  "login": "Andrey",  
  "password": "secret123"  
}  
Ответ приходит в хедере "Authorization_token"  
