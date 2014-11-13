FROM golang

RUN mkdir -p /go/src/github.com/code9io/go-angular
ADD . /go/src/github.com/code9io/go-angular

WORKDIR /go/src/github.com/code9io/go-angular

RUN go get -t ./...
RUN go build
RUN go install

EXPOSE 80
