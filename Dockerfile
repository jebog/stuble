FROM golang:alpine AS build-env
LABEL MAINTAINER = "Jean-Marc Amon <jebog@marclabs.com>"
ENV GOPATH /go
WORKDIR /go/src/app
COPY src /go/src/stuble
RUN cd /go/src/stuble && go get -v ./... && go build .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/stuble/stuble /app
COPY .env /app

EXPOSE 8000

ENTRYPOINT [ "./stuble" ]