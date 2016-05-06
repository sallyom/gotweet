FROM golang:1.7
EXPOSE 8080
ADD . /go/src/github.com/sallyom/gotweet

RUN go get github.com/ChimeraCoder/anaconda
RUN go install github.com/sallyom/gotweet/cmd/gotweet
WORKDIR /go/src/github.com/sallyom/gotweet
ENTRYPOINT ["gotweet"]
