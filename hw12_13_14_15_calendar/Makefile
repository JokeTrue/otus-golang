build:
	CGO_ENABLED=0 GOOS=linux go build -o ./.bin/calendar ./cmd/calendar/main.go

install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

test:
	go test -race -v ./...

lint: install-deps
	golangci-lint --config=./.golangcilint.yaml run ./...

.PHONY: build

