run:
	go run cmd/main.go
proto:
	protoc -I proto ./proto/monitor.proto --go_out=./proto/gen/ --go_opt=paths=source_relative --go-grpc_out=./proto/gen
startServe:
	cd server && go run ./cmd/main.go
startClient:
	cd ./client && go run ./cmd/main.go