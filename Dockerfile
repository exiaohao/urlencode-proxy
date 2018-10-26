FROM alpine:3.8

ADD dist/proxy /usr/local/bin

ENTRYPOINT ["/usr/local/bin/proxy", "&"]
