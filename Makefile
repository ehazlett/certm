CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
TAG=${TAG:-latest}
OS="darwin windows linux"
ARCH="amd64 386"
COMMIT=`git rev-parse --short HEAD`

all: deps build

deps:
	@godep restore

clean:
	@rm -rf certm certm_*

build: deps
	@godep go build -a -tags 'netgo' -ldflags "-w -X github.com/ehazlett/certm/version.GitCommit $(COMMIT) -linkmode external -extldflags -static" .

build-cross: deps
	@gox -os=$(OS) -arch=$(ARCH) -ldflags "-w -X github.com/ehazlett/certm/version.GitCommit $(COMMIT)" -output="certm_{{.OS}}_{{.Arch}}"

image: build
	@echo Building image $(TAG)
	@docker build -t ehazlett/certm:$(TAG) .
    @echo Building image $(TAG)-alpine
    @docker build -f ./Dockerfile-alpine -t ehazlett/certm:$(TAG)-alpine .
release: deps build image
	@docker push ehazlett/certm:$(TAG)
	@docker push ehazlett/certm:$(TAG)-alpine

test:
	@bats test/integration/cli.bats test/integration/certs.bats

.PHONY: all deps build clean image test release
