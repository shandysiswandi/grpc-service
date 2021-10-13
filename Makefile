proto:
	@protoc --go_out=plugins=grpc:pb/. pb/todo.proto

run:
	@go run cmd/main.go