FROM golang

ENV PORT "8080"
ADD ./bloodlines /go/bin/bloodlines
ADD ./config-dev.json /go/bin/config.json

ENTRYPOINT /go/bin/bloodlines

EXPOSE 8080
