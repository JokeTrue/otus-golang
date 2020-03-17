#!/usr/bin/env bash
set -euo pipefail

go build -o go-validate govalidate/*
./go-validate models/models.go
go test -v -tags generation ./models

rm -f go-validate
echo "PASS"
