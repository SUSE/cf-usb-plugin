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

# uploadDriver operation will is not correctly generated, ignore it for now
swagger generate client -f $GOPATH/src/github.com/hpcloud/cf-usb/swagger-spec/api.json  -t ./  -A usb-client \
 -O updateCatalog \
 -O getInfo \
 -O getInstances \
 -O createInstance \
 -O getInstance \
 -O updateInstance \
 -O deleteInstance \
 -O pingInstance \
 -O getServiceByInstanceId \
 -O getService \
 -O updateService \
 -O getServicePlans \
 -O getServicePlan \
 -O updateServicePlan \
 -O getAllDials \
 -O createDial \
 -O getDial \
 -O updateDial \
 -O deleteDial \
 -O getDriver \
 -O uploadDriver

popd
