FROM scratch
ADD certm /bin/certm
ENTRYPOINT ["/bin/certm"]
