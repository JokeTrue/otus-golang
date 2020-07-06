FROM golang:1.14-alpine AS build

WORKDIR $GOPATH/src/github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar

COPY . ./
RUN go mod download
RUN go build -o /go/bin/calendar_scheduler ./cmd/calendar_scheduler/main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/calendar_scheduler /bin/