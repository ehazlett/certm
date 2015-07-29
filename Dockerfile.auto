FROM golang:1.4-cross

RUN go get github.com/tools/godep
RUN go get github.com/aktau/github-release
RUN go get github.com/mitchellh/gox

RUN gox -build-toolchain

ADD . /go/src/github.com/ehazlett/certm

ENV BINARY_SHA 386f6e91114dc252a13b266fe2ac3a27e83bd0f7
RUN curl -fL https://get.docker.com/builds/Linux/x86_64/docker-1.7.1 -o /usr/local/bin/docker \
  && echo "$BINARY_SHA /usr/local/bin/docker" | sha1sum -c -

RUN chmod +x /usr/local/bin/docker

RUN (git clone https://github.com/sstephenson/bats.git && \
    cd bats && ./install.sh /usr/local)

ENV SHELL /bin/bash

WORKDIR /go/src/github.com/ehazlett/certm
RUN make && cp certm /bin/certm
ENTRYPOINT ["/bin/certm"]
