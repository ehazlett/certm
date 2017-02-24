CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
APP=certm
REPO=ehazlett/$(APP)
TAG=${TAG:-latest}
OS="darwin windows linux"
ARCH="amd64 386"
COMMIT=`git rev-parse --short HEAD`

all: build

clean:
	@rm -rf certm certm_*

build:
	@go build -a -tags 'netgo' -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT) -linkmode external -extldflags -static" .

build-cross:
	@gox -os=$(OS) -arch=$(ARCH) -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT) -linkmode external -extldflags -static" -output="certm_{{.OS}}_{{.Arch}}"

image: build
	@echo Building image $(TAG)
	@docker build -t $(REPO):$(TAG) .

release: deps build image
	@docker push $(REPO):$(TAG)

test:
	@bats test/integration/cli.bats test/integration/certs.bats

.PHONY: all build clean image test release
