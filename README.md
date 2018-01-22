# cf-plugin-usb

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

### cf usb target

Before you can use the plugin, a usb management api must be set, to do this the following command must be executed:

**Usage**:

```bash
cf usb target {usb_management_endpoint}
```

### cf usb info

Displays the information about the broker API version and the Universal Service Broker version

**Usage**:

```bash
cf usb info
```

### cf usb create-instance

Creates a drivers instance.

**Arguments**:
- driverName:
  - Required: true
  - Description: The type of the driver. Example: mysql, mssql, etc.
- instanceName:
  - Required: true
  - Description: The name of the new driver instance. This will also be the default name of the Cloud Foundry Service
- configFile:
  - Required: false
  - Description: a configuration file for the driver instance
- configValue:
  - Required: false
  - Description: the configuration of the driver instance in JSON format

If the <configFile> and <configValue> arguments are not specified, the plugin will generate a wizard based on the driver config schema.

**Usage**:

```bash
cf usb create-instance <driverName> <instanceName> configValue/configFile <jsonValue/filePath>
```

### cf usb delete-instance

Deletes a driver instance.

**Arguments**
- instanceName
  - Required: true
  - Description: The driver instance name.

**Usage**

```bash
cf usb delete-instance <instanceName>
```  

### cf usb create-driver

Creates a driver

**Arguments**
- driverType
  - Required: true
  - Description: The type of the driver
- driverName
  - Required: true
  - Description: The name of the driver
- driverPath
  - Required: true
  - Description: Path to the driver binaries

**Usage**

```bash
cf usb create-driver <driverType> <driverName> <driverPath>
```

### cf usb update-driver

Updates the driver name.

**Arguments**
- oldDriverName
  - Required: true
  - Description: The name of the driver that is going to be update-driver
- newDriverName
  - Required: true
  - Description: The new name of the driver

```bash
cf usb update-driver <oldDriverName> <newDriverName>
```

### cf usb update-instance

Updates the configuration for a driver instance.

**Argument**
- instanceName
  - Required: true
  - Description: the name of the driver instance that is going to be updated
  - configFile:
    - Required: false
    - Description: a configuration file for the driver instance
  - configValue:
    - Required: false
    - Description: the configuration of the driver instance in JSON format

If the <configFile> and <configValue> arguments are not specified, the plugin will generate a wizard based on the driver config schema.

**Usage**

```bash
cf usb update-instance <instanceName> configValue/configFile <jsonValue/filePath>
```

### cf usb update-service

Updates the catalog information for the exposed Cloud Foundry Service

**Arguments**
- instanceName
  - Required: true
  - Description: the name of the Driver instance

**Usage**

```bash
cf usb update-service <instanceName>
```

After executing a command, a wizard is displayed that allows the user to change the catalog information for a service

### cf usb delete-driver

Deletes and existing driver

**Arguments**

- driverName
  - Required: true
  - Description: the name of the driver that will be deleted

**Usage**

```bash
cf usb delete-driver <driverName>
```

### cf usb drivers

Lists all the available drivers

**Usage**

```bash
cf usb drivers
```

### cf usb instances

Lists all the driver instances for a driver

**Arguments**

- driverName
  - Required: true
  - Description: the name of the driver

**Usage**

```bash
cf usb instances <driverName>
```

### cf usb dials

Lists the dials for a driver instance

**Arguments**
- instanceName
  - Required: true
  - Description: The name of the driver instance

**Usage**

```bash
cf usb dials <driverInstance>
```
