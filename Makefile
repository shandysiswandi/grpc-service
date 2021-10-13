proto:
	@protoc --go_out=plugins=grpc:pb/. pb/todo.proto

run:
	@go run cmd/main.go

test:
	@go test ./... -coverprofile coverage.out

sonar:
	@sonar-scanner \
  	-Dsonar.projectKey=grpc-service \
  	-Dsonar.sources=. \
  	-Dsonar.host.url=http://localhost:9000 \
  	-Dsonar.login=56c30d1f2fc7ba5d6788b1d06e39b1f0f8e60707