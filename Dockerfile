FROM golang:1.12.10

ENV GOPATH=/go
RUN export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
ENV GO111MODULE=off
RUN mkdir -p /go/src/sozluk
ADD . /go/src/sozluk

WORKDIR /go/src/sozluk
CMD cd /conf && [[ ! -e app.conf ]] && cp app.conf.dist app.conf

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

EXPOSE 9000
ENTRYPOINT revel run /go/src/sozluk 9000
