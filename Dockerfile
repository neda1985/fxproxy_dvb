FROM golang:1.18-alpine as builder

ADD . /go/src/fxproxy

RUN cd /go/src/fxproxy && go build .
RUN go install github.com/traefik/whoami@latest

FROM alpine:3.6

COPY --from=builder /go/src/fxproxy/fxproxy /fxproxy
COPY --from=builder /go/bin/whoami /whoami
ADD run.sh /run.sh
RUN chmod a+x /run.sh

ENTRYPOINT ["/run.sh"]

EXPOSE 8080