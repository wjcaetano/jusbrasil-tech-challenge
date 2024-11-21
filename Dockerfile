FROM golang:alpine3.19

ENV APPLICATION_PACKAGE=./cmd/api
ENV APP_PATH=/app
ENV GO_TEST_FLAGS="-tags=integration -p=1"
ENV SCOPE=local

RUN apk update && apk add --no-cache bash nodejs npm mysql-client

RUN	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.62.0/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.62.0

WORKDIR $APP_PATH

RUN go install github.com/arthurgustin/godepth@latest
RUN go install github.com/arch-go/arch-go@latest
RUN npm -g install directory-validator

ADD ./commands/*.sh /commands/
ADD .golangci-project.yml /commands/
ADD arch-go.yml /commands/
ADD .directoryvalidator.json /commands/
RUN chmod a+x /commands/*.sh
ADD ./resources /resources
ADD . .

RUN go mod tidy
RUN go build -o main $APPLICATION_PACKAGE

RUN chmod +x /commands/*.sh

ENV GO_RACE_DETECTION_ENABLED="true"