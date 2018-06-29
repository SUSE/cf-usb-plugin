# cf-usb-plugin

Universal Service Broker plugin for Cloud Foundry CLI.

It is part of the [CF-USB project](https://github.com/SUSE/cf-usb).

To use the plugin you must have *Admin* Cloud Foundry privileges.

## Building

For your first build:

```bash
mkdir -p $GOPATH/src/github.com/SUSE
git clone https://github.com/SUSE/cf-usb-plugin $GOPATH/src/github.com/SUSE
cd $GOPATH/src/github.com/SUSE/cf-usb-plugin
make tools
make
```

The build artifacts will be placed in *bin/{os}/{arch}/*

## Installing

The cf cli usb plugin requires cf cli version 6.14 or greater

After building the plugin, you can install it using the following command:

```bash
cf install-plugin <path_to_the_artifact>
```

To uninstall:

```bash
cf uninstall-plugin usb
```

## Testing

To run the cf cli usb plugin tests, run the following command:

```bash
cd $GOPATH/src/github.com/SUSE/cf-usb-plugin
make test
```

## Commands

### cf usb-target

Before you can use the plugin, a usb management api must be set, to do this the following command must be executed:

**Usage**:

```bash
cf usb-target {usb_management_endpoint}
```

### cf usb-info

Displays the information about the broker API version and the Universal Service Broker version

**Usage**:

```bash
cf usb-info
```

### cf usb-create-driver-endpoint

Creates a driver endpoint

**Arguments**
- Name
  - Required: true
  - Description: The name of the driver endpoint
- Endpoint URL
  - Required: true
  - Description: The URL
- Authentication Key
  - Required: true
  - Description: Path to the driver binaries with an option to provide a JSON file for the metadata in a format of 'mkey1:mval1;mkey2:mval2'

**Usage**

```bash
cf usb-create-driver-endpoint NAME ENDPOINT_URL AUTHENTICATION_KEY [-c METADATA]
```

### cf usb-update-driver-endpoint

Updates the configuration for a driver instance.

**Argument**
- Name
  - Required: true
  - Description: the name of the driver instance that is going to be updated
- Endpoint URL
  - Required: true
  - Description: The URL
- Authentication Key
  - Required: true
  - Description: Path to the driver binaries with an option to provide a file with the metadata as JSON

**Usage**

```bash
cf usb-update-driver-endpoint NAME [-t ENDPOINT_URL] [-k AUTHENTICATION_KEY] [-c METADATA_AS_JSON]
```

### cf usb-delete-driver-endpoint

Deletes an existing driver

**Arguments**

- Name
  - Required: true
  - Description: the name of the driver that will be deleted

**Usage**

```bash
cf usb-delete-driver-endpoint <driverName>
```

### cf usb-driver-endpoints

Lists all the available drivers

**Usage**

```bash
cf usb-driver-endpoints
```

