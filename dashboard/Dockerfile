FROM golang:1.20 as build

ENV http_proxy  ""
ENV https_proxy ""

WORKDIR /go/src/github.com/horsing/goflow/dashboard

ADD . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o handler .

CMD ["./handler"]
