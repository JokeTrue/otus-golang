build:
	CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

install-deps:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0

run: build
	docker-compose up --build server

test:
	go test -race -v ./...

lint: install-deps
	golangci-lint --config=./.golangcilint.yaml run ./...

.PHONY: build
