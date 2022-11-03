FROM golang:1.17-alpine3.13

ENV CGO_ENABLED=0
RUN mkdir /app
WORKDIR /app
ADD go.mod go.sum /app
ADD *.go /app
RUN go test .
