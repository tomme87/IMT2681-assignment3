FROM golang:1.9.2-stretch

LABEL maintainer="tomrf@stud.ntnu.no"

RUN apt-get update && apt-get -y upgrade

RUN go get github.com/tomme87/IMT2681-assignment3/cmd/webapp

RUN go install github.com/tomme87/IMT2681-assignment3/cmd/webapp

WORKDIR $GOPATH

CMD ["./bin/webapp"]
