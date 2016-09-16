#!/usr/bin/env bash

set -v
set -e

go get -v github.com/go-swagger/go-swagger/cmd/swagger

rm -rf $GOPATH/src/github.com/hpcloud/cf-plugin-usb/lib/models
rm -rf $GOPATH/src/github.com/hpcloud/cf-plugin-usb/lib/client

pushd $GOPATH/src/github.com/hpcloud/cf-plugin-usb
	mkdir -p lib
popd

pushd $GOPATH/src/github.com/hpcloud/cf-plugin-usb/lib

../.tools/swagger generate client -f $GOPATH/src/github.com/hpcloud/cf-usb/swagger-spec/management-api.json  -t ./  -A usb-client 
popd
