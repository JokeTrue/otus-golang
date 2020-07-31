build:
	go build -o ./bin/calendar ./cmd/calendar/main.go
	go build -o ./bin/calendar_sender ./cmd/calendar_sender/main.go
	go build -o ./bin/calendar_scheduler ./cmd/calendar_scheduler/main.go

.PHONY: test
test:
	go test -race -v ./pkg/...

.PHONY: lint
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint --config=./.golangcilint.yaml run ./...

.PHONY: generate
generate:
	apt install -y protobuf-compiler
	protoc --proto_path=./event/delivery/grpc/schema --go_opt=paths=source_relative --go_out=plugins=grpc:./event/delivery/grpc/schema/ event.proto

up:
	docker-compose -f ./docker/docker-compose.yml up

down:
	docker-compose -f ./docker/docker-compose.yml down

bdd:
	docker-compose -f ./docker/docker-compose-tests.yml up --abort-on-container-exit --exit-code-from integration-tests && \
	docker-compose -f ./docker/docker-compose-tests.yml down