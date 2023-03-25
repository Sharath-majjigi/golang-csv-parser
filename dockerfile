# Build Stage

FROM golang:1.17-alpine as build-env

ENV APP_NAME golang-csv-parser
ENV CMD_PATH main.go

COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

RUN CGO_ENABLED=O go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# Run Stage

FROM alpine:3.14

ENV APP_NAME golang-csv-parser

COPY --from=build-env /$APP_NAME .

EXPOSE 8002

# start app
CMD ./$APP_NAME

