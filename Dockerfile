FROM golang:1.17-alpine3.13

ENV CGO_ENABLED=0
RUN go install github.com/pion/webrtc/v3/examples/pion-to-pion/offer@latest
RUN go install github.com/pion/webrtc/v3/examples/pion-to-pion/answer@latest

RUN mkdir /app
WORKDIR /app
ADD go.mod go.sum /app/
ADD *.go /app/
RUN go test -timeout 3s .
