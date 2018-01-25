#!/usr/bin/env bash

set -v
set -e

go get -v github.com/go-swagger/go-swagger/cmd/swagger

rm -rf $GOPATH/src/github.com/SUSE/cf-usb-plugin/lib/models
rm -rf $GOPATH/src/github.com/SUSE/cf-usb-plugin/lib/client

pushd $GOPATH/src/github.com/SUSE/cf-usb-plugin
	mkdir -p lib
popd

pushd $GOPATH/src/github.com/SUSE/cf-usb-plugin/lib

swagger generate client -f $GOPATH/src/github.com/SUSE/cf-usb/swagger-spec/management-api.json  -t ./  -A usb-client 
popd
