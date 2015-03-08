CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
TAG=${TAG:-latest}

all: deps build

deps:
	@godep restore

clean:
	@rm -rf Godeps/_workspace cert-tool

build: deps
	@godep go build -a -tags 'netgo' -ldflags '-w -linkmode external -extldflags -static' .

image: build
	@echo Building image $(TAG)
	@docker build -t ehazlett/cert-tool:$(TAG) .

release: deps build image
	@docker push ehazlett/cert-tool:$(TAG)

test:
	@bats test/integration/cli.bats test/integration/certs.bats

.PHONY: all deps build clean image test release
