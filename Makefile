run:
	go run cmd/main.go
proto:
	protoc -I proto proto/monitor.proto --go_out=./gen/ --go_opt=paths=source_relative --go-grpc_out=./gen