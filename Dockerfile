FROM golang:1.8.1-alpine

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV PROJECT $GOPATH/src/github.com/swatlabs/GoDataberus

RUN mkdir -p /go/src/github.com/swatlabs/GoDataberus
RUN mkdir -p "$GOPATH/bin"
RUN chmod -R 777 "$GOPATH"

RUN apk add --update git && rm -rf /var/cache/apk/*
RUN go get github.com/Masterminds/glide
RUN go install github.com/Masterminds/glide

WORKDIR $PROJECT

VOLUME $PROJECT

ADD ./ $PROJECT
RUN glide install
RUN go install

ENTRYPOINT $GOPATH/bin/GoDataberus

EXPOSE 8080
