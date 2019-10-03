FROM golang:1.11.5

RUN export GOPATH=/go
RUN export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
ENV GO111MODULE=on
RUN mkdir -p /go/src/sozluk
ADD . /go/src/sozluk

WORKDIR /go/src/sozluk
CMD cd /conf && [[ ! -e app.conf ]] && cp app.conf.dist app.conf

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN dep ensure  --vendor-only

EXPOSE 9000
ENTRYPOINT revel run /go/src/sozluk 9000
