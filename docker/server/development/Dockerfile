FROM golang:1.16

ARG WORKING_DIR
WORKDIR $WORKING_DIR

RUN go install github.com/go-delve/delve/cmd/dlv@latest

ENV GO111MODULE=on
ENV GOPATH=''