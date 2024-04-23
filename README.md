Backend for matching application



==DEV==
protoc --go_out=../gen/go --go_opt=paths=source_relative --go-grpc_out=../gen/go --go-grpc_opt=paths=source_relative protobuf_services/auth/auth.proto
migrate -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable -path migrations up
migrate -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable -path auth_service/migrations up
./api_gateway_service/bin/main && ./auth_service/bin/main &&
=========asd
version: '3.8'
  
services:
  unimatch:
    build: ./
    command: ./api_gateway_service/bin/main && ./auth_service/bin/main && migrate -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable -path auth_service/migrations up
    ports:
      - 8080:8080
    depends_on:
      - postgres_dbauth
    environment:
      - DB_PASSWORD=secret
      - CONFIG_GATEWAY_PATH=./api_gateway_service/configs/config.yaml
      - CONFIG_AUTH_PATH=./auth_service/configs/config.yaml
  postgres_dbauth:
    restart: always
    image: postgres:latest
    ports:
      - 5432:5432
    environment:
      - POSTGRES_PASSWORD=secret
      - POSTGRES_DBNAME=postgres
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    container_name: zoo
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_SERVER_ID: 1  
    ports:
      - "2181:2181"
  kafka-1:
    image: confluentinc/cp-kafka:latest
    container_name: broker-1
    ports:
      - "9092:9092"
      - "29092:29092"
    environment:
      KAFKA_ADVERTISED_LISTENERS: INTERNAL://kafka-1:19092,EXTERNAL://${DOCKER_HOST_IP:-127.0.0.1}:9092,DOCKER://host.docker.internal:29092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: INTERNAL:PLAINTEXT,EXTERNAL:PLAINTEXT,DOCKER:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: INTERNAL
      KAFKA_ZOOKEEPER_CONNECT: "zookeeper:2181"
      KAFKA_BROKER_ID: 1
    depends_on:
      - zookeeper
  kafka-2:
    image: confluentinc/cp-kafka:latest
    container_name: broker-2
    ports:
      - "9093:9093"
      - "29093:29093"
    environment:
      KAFKA_ADVERTISED_LISTENERS: INTERNAL://kafka-2:19093,EXTERNAL://${DOCKER_HOST_IP:-127.0.0.1}:9093,DOCKER://host.docker.internal:29093
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: INTERNAL:PLAINTEXT,EXTERNAL:PLAINTEXT,DOCKER:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: INTERNAL
      KAFKA_ZOOKEEPER_CONNECT: "zookeeper:2181"
      KAFKA_BROKER_ID: 2
    depends_on:
      - zookeeper
  kafka-3:
    image: confluentinc/cp-kafka:latest
    container_name: broker-3
    ports:
      - "9094:9094"
      - "29094:29094"
    environment:
      KAFKA_ADVERTISED_LISTENERS: INTERNAL://kafka-3:19094,EXTERNAL://${DOCKER_HOST_IP:-127.0.0.1}:9094,DOCKER://host.docker.internal:29094
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: INTERNAL:PLAINTEXT,EXTERNAL:PLAINTEXT,DOCKER:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: INTERNAL
      KAFKA_ZOOKEEPER_CONNECT: "zookeeper:2181"
      KAFKA_BROKER_ID: 3
    depends_on:
      - zookeeper
=====================================sdc
version: '3.8'
  
services:
  api_gateway_service:
    container_name: api_gateway_service
    build: ./
    command: migrate -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable -path auth_service/migrations up
    ports:
      - 8080:8080
      - CONFIG_GATEWAY_PATH=./api_gateway_service/configs/config.yaml
      
  auth_service:
    container_name: auth_service
    depends_on:
      - postgres_dbauth
    environment:
      - DB_PASSWORD=secret
      - CONFIG_AUTH_PATH=./auth_service/configs/config.yaml
  postgres_dbauth:
    restart: always
    image: postgres:latest
    ports:
      - 5432:5432
    environment:
      - POSTGRES_PASSWORD=secret
      - POSTGRES_DBNAME=postgres
    