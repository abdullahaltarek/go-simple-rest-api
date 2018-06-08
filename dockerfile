FROM golang

ADD . /go/src/github.com/cyantarek/golang-simple-api

RUN go install /go/src/github.com/cyantarek/golang-simple-api

ENTRYPOINT /go/bin/golang-simple-api

EXPOSE 8075