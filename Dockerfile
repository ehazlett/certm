FROM scratch
ADD cert-tool /bin/cert-tool
ENTRYPOINT ["/bin/cert-tool"]
