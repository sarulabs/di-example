FROM golang:1.10-alpine

WORKDIR /go/src/github.com/sarulabs/di-example

RUN apk add --no-cache git \
    && go get -u golang.org/x/vgo

COPY . .

RUN vgo build -o /go/bin/di_example

ENV SERVER_PORT=8080
ENV MONGO_URL=mongo:27017

CMD /go/bin/di_example