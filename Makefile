include version.mk

ARCH:=$(shell go env GOOS).$(shell go env GOARCH)
COMMIT_HASH=$(shell git log --pretty=format:'%h' -n 1)
APP_VERSION=$(VERSION)-$(COMMIT_HASH)

PKGSDIRS=$(shell go list -f '{{.ImportPath}}' ./... | grep -v /vendor/)

IMAGE_NAME=splatform/cf-usb-plugin
IMAGE_TAG=$(subst +,_,$(APP_VERSION))

print_status = @printf "\033[32;01m==> $(1)\033[0m\n"

.PHONY: all clean format lint vet bindata build test docker docker-write-tag-files docker-push
all: clean format lint vet bindata test build

clean:
	$(call print_status, Cleaning)
	rm -f cf-usb-plugin
	rm -f cf-usb-plugin-*.tgz
	rm -f cf-usb-plugin-*.tar
	rm -f .cf-usb-plugin-docker-*.txt

format:
	$(call print_status, Checking format)
	@echo $(PKGSDIRS) | tr ' ' '\n' | xargs -I '{p}' -n1 goimports -e -l {p} | sed "s/^/Failed: /"

lint:
	$(call print_status, Linting)
	@echo $(PKGSDIRS) | tr ' ' '\n' | xargs -I '{p}' -n1 golint {p}| grep -v "fakes/*" | grep -v "lib/*" | grep -v "mock_.*\.go" | sed "s/^/Failed: /"

vet:
	$(call print_status, Vetting)
	go vet $(PKGSDIRS)

build:
	$(call print_status, Building)
	$(call gobuild,linux,amd64)
	$(call gobuild,windows,amd64,.exe)
	$(call gobuild,darwin,amd64)

gobuild = GOARCH=$(2) GOOS=$(1) go build \
		-ldflags="-X main.version=$(APP_VERSION)" \
		-o="build/$(1)-$(2)/cf-usb-plugin$(3)" ./

linux_dist: build
	$(call print_status, Disting linux)
	$(call godist,linux,amd64)

windows_dist: build
	$(call print_status, Disting windows)
	$(call godist,windows,amd64)

darwin_dist: build
	$(call print_status, Disting darwin)
	$(call godist,darwin,amd64)

dist: linux_dist \
	windows_dist \
	darwin_dist

godist = GOARCH=$(2) GOOS=$(1) tar czf cf-usb-plugin-$(APP_VERSION)-$(1)-$(2).tgz build/$(1)-$(2)/*

tools:
	$(call print_status, Installing Tools)
	@case $$(go version) in \
         "go version go1.[1-5]*") go get -u golang.org/x/tools/cmd/vet ;; \
	esac
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/golang/lint/golint

test:
	$(call print_status, Testing)
	go test -cover $(PKGSDIRS)

docker:
	$(call print_status, Creating docker image)
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	docker tag -f $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):$(BRANCH)

# Used by ci
docker-write-tag-files:
	$(call print_status, Writing docker tag files)
	echo $(IMAGE_TAG) > cf-usb-plugin-docker-version-tag.txt
	echo $(BRANCH) > cf-usb-plugin-docker-branch-tag.txt

docker-push: docker
	$(call print_status, Pushing docker image)
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(IMAGE_NAME):$(BRANCH)
