EXECUTABLE := main

build:
	@go build -o $(EXECUTABLE) .

run: build
	@./$(EXECUTABLE)

test:
	@go test -v ./...

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: proto