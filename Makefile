proto:
	protoc -I proto proto/monitor.proto --go_out=./gen/ --go_opt=paths=s    ource_relative --go-grpc_out=./gen