FROM golang

ADD ./go/ .
RUN go install main
ENTRYPOINT /go/bin/main

EXPOSE 8080
