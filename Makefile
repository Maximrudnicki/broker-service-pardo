
build:
	protoc -Iproto --go_opt=module=broker-service --go_out=. --go-grpc_opt=module=broker-service --go-grpc_out=. proto/*.proto
	go build -o bin/broker-service.exe ./cmd/.