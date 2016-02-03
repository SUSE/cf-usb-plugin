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
 -O getDrivers \
 -O createDriver \
 -O getDriver \
 -O updateDriver \
 -O deleteDriver \
 -O getDriverSchema \
 -O getDialSchema \
 -O getDriverInstances \
 -O createDriverInstance \
 -O getDriverInstance \
 -O updateDriverInstance \
 -O deleteDriverInstance \
 -O pingDriverInstance \
 -O getServiceByInstanceId \
 -O getService \
 -O updateService \
 -O getServicePlans \
 -O createServicePlan \
 -O getServicePlan \
 -O updateServicePlan \
 -O deleteServicePlan \
 -O getAllDials \
 -O createDial \
 -O getDial \
 -O updateDial \
 -O deleteDial \
 -O getDriver

popd
