build:
	go build -o ./bin/calendar ./cmd/calendar/main.go
	go build -o ./bin/calendar_sender ./cmd/calendar_sender/main.go
	go build -o ./bin/calendar_scheduler ./cmd/calendar_scheduler/main.go

install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	apt install -y protobuf-compiler

test:
	go test -race -v ./...

lint: install-deps
	golangci-lint --config=./.golangcilint.yaml run ./...

generate:
	protoc --proto_path=./event/delivery/grpc/schema --go_opt=paths=source_relative --go_out=plugins=grpc:./event/delivery/grpc/schema/ event.proto

.PHONY: build

