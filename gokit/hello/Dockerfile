FROM golang

RUN mkdir -p /go/src/suiyidian.cn/sonhineboy/gokit

ADD . /go/src/suiyidian.cn/sonhineboy/gokit

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run suiyidian.cn/sonhineboy/gokit/hello/cmd  -watch suiyidian.cn/sonhineboy/gokit/hello
