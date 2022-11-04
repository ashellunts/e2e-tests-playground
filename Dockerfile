FROM golang:1.17-alpine3.13

ENV CGO_ENABLED=0

RUN mkdir /test
WORKDIR /test
ADD go.mod go.sum /test/
RUN go mod download

COPY offer /test/offer
COPY answer /test/answer
RUN go build -o offer_app ./offer
RUN go build -o answer_app ./answer

ADD main_test.go /test/
CMD go test -v -timeout 20s .
