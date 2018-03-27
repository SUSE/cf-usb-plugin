GIT_DESCRIBE ?= $(shell git describe --tags --long)

GIT_TAG ?= $(shell echo $(GIT_DESCRIBE) | awk -F - '{ print $$1 }' )
# GIT_COMMITS will be 0 if we're on the tag
GIT_COMMITS ?= $(shell echo $(GIT_DESCRIBE) | awk -F - '{ print $$2 }' )
GIT_SHA ?= $(shell echo $(GIT_DESCRIBE) | awk -F - '{ print $$3 }' )
APP_VERSION ?= $(GIT_TAG)+$(GIT_COMMITS).$(GIT_SHA)
