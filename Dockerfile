FROM scratch
COPY certm /bin/certm
ENTRYPOINT ["/bin/certm"]
