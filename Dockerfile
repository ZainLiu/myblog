FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/ZainLiu_github.com/lzyblog
COPY . $GOPATH/src/ZainLiu_github.com/lzyblog
RUN go build .

EXPOSE 8088
ENTRYPOINT ["./lzyblog"]