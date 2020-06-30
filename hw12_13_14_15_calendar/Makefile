build:
	go build -o ./bin/calendar ./cmd/calendar/main.go

test:
	go test -race -v ./...

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint --config=./.golangcilint.yaml run ./...

generate:
	protoc --proto_path=./event/delivery/grpc/schema --go_opt=paths=source_relative --go_out=plugins=grpc:./event/delivery/grpc/schema/ event.proto

.PHONY: build

