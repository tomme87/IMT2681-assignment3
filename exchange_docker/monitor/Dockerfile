FROM golang:1.9.2-stretch

LABEL maintainer="tomrf@stud.ntnu.no"

RUN apt-get update && apt-get -y upgrade

RUN go get github.com/tomme87/IMT2681-assignment2/cmd/monitor

RUN go install github.com/tomme87/IMT2681-assignment2/cmd/monitor

WORKDIR $GOPATH

CMD ["./bin/monitor"]
