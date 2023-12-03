run:
	@echo "Выполняется запуск приложения."
	go run cmd/main.go
proto:
	@echo "Генерация скриптов protobuf."
	protoc -I proto ./proto/monitor.proto --go_out=./proto/gen/ --go_opt=paths=source_relative --go-grpc_out=./proto/gen

.PHONY: startServe
startServe:
	@echo "Выполняется запуск сервера."
	cd server && go run cmd/main.go

.PHONY: startClient
startClient:
	@echo "Выполняется запуск клиента."
	cd client && go run cmd/main.go

.PHONY: tests
tests:
	@echo "Выполняется запуск тестов."
	cd server && go test -v -cover ./...
