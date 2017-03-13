FROM alpine:latest
COPY certm /usr/local/bin/certm
ENTRYPOINT ["certm"]