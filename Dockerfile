FROM golang:1.12

RUN mkdir -p $GOPATH/src/two
WORKDIR $GOPATH/src/two

COPY . .
RUN cd $GOPATH/src/two/ && go get

EXPOSE 7778

RUN go get
RUN go build

CMD ["./two"]
