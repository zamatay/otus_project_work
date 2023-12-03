ROOT_DIR        := $(shell pwd)
TEST_DIR        := $(ROOT_DIR)/_test
MAIN_DIR        := $(ROOT_DIR)/cmd/app
SERVER_DIR        := $(ROOT_DIR)/server/cmd
CLIENT_DIR        := $(ROOT_DIR)/client/cmd

run:
	@echo "Выполняется запуск приложения."
	go run cmd/main.go
proto:
	@echo "Генерация скриптов protobuf."
	protoc -I proto ./proto/monitor.proto --go_out=./proto/gen/ --go_opt=paths=source_relative --go-grpc_out=./proto/gen

.PHONY: startServe
startServe:
	@echo "Выполняется запуск сервера."
	go run $(SERVER_DIR)/main.go

.PHONY: startClient
startClient:
	@echo "Выполняется запуск клиента."
	go run $(CLIENT_DIR)/main.go

.PHONY: tests
tests:
	@echo "Выполняется запуск тестов."
	CGO_ENABLED=1 $(GO) test -json -race