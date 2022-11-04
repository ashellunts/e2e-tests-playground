FROM golang:1.17-alpine3.13

ENV CGO_ENABLED=0
RUN go install github.com/pion/webrtc/v3/examples/pion-to-pion/offer@latest
RUN go install github.com/pion/webrtc/v3/examples/pion-to-pion/answer@latest

RUN mkdir /app
WORKDIR /app
ADD go.mod go.sum /app/
COPY offer /app/offer
COPY answer /app/answer
ADD *.go /app/
RUN go build -o offer_app ./offer
RUN go build -o answer_app ./answer

RUN go test -v -timeout 10s .
