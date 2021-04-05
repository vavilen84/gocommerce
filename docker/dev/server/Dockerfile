FROM golang:1.16

ARG WORKDIR
WORKDIR $WORKDIR
ENV GO111MODULE=on
ENV GOPATH=''
COPY ./server/ ./

CMD go mod vendor && go mod tidy && go run .
