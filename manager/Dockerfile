FROM golang:1.8

COPY vendor/* /go/src/

WORKDIR /go/src/github.com/whobrokethebuild/manager
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80

CMD ["manager"]
