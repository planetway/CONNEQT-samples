FROM alpine:3.4

RUN apk --update upgrade && \
    apk add bash ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ADD bin/go-time /go-time

WORKDIR /

CMD ["/go-time"]
