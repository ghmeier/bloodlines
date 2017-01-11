FROM golang

ENV PORT "8080"
#ADD . /go/src/github.com/ghmeier/bloodlines/
ADD ./bloodlines /go/bin/bloodlines
ADD ./config.json /go/bin/config.json
#RUN go install github.com/ghmeier/bloodlines

ENTRYPOINT /go/bin/bloodlines

EXPOSE 8080
