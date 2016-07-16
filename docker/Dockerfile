FROM golang:1.6-alpine

ENV APP_DIR $GOPATH/src/github.com/kieranajp/gourl

ADD . $APP_DIR
WORKDIR $APP_DIR

RUN apk --update add git && rm -rf /var/cache/apk/*

RUN go get -u github.com/codegangsta/gin
RUN go get -u github.com/tools/godep
RUN godep restore

EXPOSE 3000
CMD gin run
