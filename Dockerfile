FROM golang

RUN apt-get update
RUN apt-get -y install zip

RUN go get github.com/GeertJohan/go.rice/rice
RUN go get github.com/tools/godep

ADD . /go/src/github.com/code9io/go-angular
WORKDIR /go/src/github.com/code9io/go-angular

RUN go get -t ./...
RUN go install

EXPOSE 80
