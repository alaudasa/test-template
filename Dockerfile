FROM golang:1.8

COPY . $GOPATH/src/github.com/alaudasa/test-template

WORKDIR $GOPATH/src/github.com/alaudasa/test-template

RUN go install

CMD ["test-template", "run"]