FROM golang:1.14-alpine AS build

RUN set -ex  && apk add --no-cache --virtual build-dependencies build-base git \
             && go get -ldflags "-extldflags -static" github.com/pressly/goose/cmd/goose \
             && apk del build-dependencies

FROM alpine:3.10
COPY --from=build /go/bin/goose /bin/

ENTRYPOINT ["goose"]