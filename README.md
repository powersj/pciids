# PCIIDs

*Lookup vendor and device names using PCI IDs!*

[![Build Status](https://travis-ci.org/powersj/pciids.svg?branch=master)](https://travis-ci.org/powersj/pciids/) [![Go Report Card](https://goreportcard.com/badge/github.com/powersj/pciids)](https://goreportcard.com/report/github.com/powersj/pciids) [![Go Reference](https://pkg.go.dev/badge/github.com/powersj/pciids.svg)](https://pkg.go.dev/github.com/powersj/pciids)

## CLI

To search for devices using the CLI, pass in either:

  a) a pair of vendor and device PCI IDs
  b) two pairs, vendor and device PCI IDs as well as sub-vendor and
     sub-device PCI IDs:

Here are some examples:

```text
$ pciids 1d0f efa1
1d0f:efa1           - Amazon.com, Inc. Elastic Fabric Adapter (EFA)
$ pciids 10de 2206 10de 1467
10de:2206 10de:1467 - NVIDIA Corporation GA102 [GeForce RTX 3080]
```

If there are multiple matches then all matches are returned.

### JSON output

The command can take the `--json` flag to produce JSON output:

```json
$ pciids 121a 0009 121a 0009 --json
[
    {
        "vendorID": "121a",
        "deviceID": "0009",
        "vendorName": "3Dfx Interactive, Inc.",
        "deviceName": "Voodoo 4 / Voodoo 5",
        "subVendorID": "121a",
        "subDeviceID": "0009",
        "subVendorName": "3Dfx Interactive, Inc.",
        "subDeviceName": "Voodoo5 AGP 5500/6000"
    }
]
```

### Debug output

The `--debug` flag to produce additional output while running:

```json
$ pciids --json 121a 0009 121a 0009
DEBU Looking up 121a:0009 121a:0009
DEBU Downloading https://raw.githubusercontent.com/pciutils/pciids/master/pci.ids
DEBU 200 OK
DEBU Parsing vendor IDs
DEBU Parsing PCI IDs
DEBU Found 1 results
121a:0009 121a:0009 - 3Dfx Interactive, Inc. Voodoo5 AGP 5500/6000
```

## Install

Below outlines the various ways to obtain and install pciids.

### From binary

Download the [latest release](https://github.com/powersj/pciids/releases/latest)
of pciids for your platform and extract the tarball:

```shell
wget pciids<version>_<os>_<arch>.tar.gz
tar zxvf pciids<version>_<os>_<arch>.tar.gz
```

The tarball will extract the readme, license, and the pre-compiled binary.

### From source

To build and install directly from source run:

```shell
git clone https://github.com/powersj/pciids
cd pciids
make
```

The default make command will run `go build` and produce a binary in the root
directory.

### From go

To download using the `go get` command run:

```shell
go get github.com/powersj/pciids
```

The executable object file location will exist at `${GOPATH}/bin/pciids`

## API usage

Users can take advantage of various functions in their own code:

* `All()`: returns a list of all PCI ID devices
* `LatestFile()`: string of the latest PCI ID database
* `Parse(string)`: parses a given PCI ID database string
* `QueryDevice(vendorID, deviceID)`: Searches for devices matching a PCI ID
   pair
* `QuerySubDevice(vendorID, deviceID, subVendorID, subDeviceID)`: Like
   QueryDevice, but matches two PCI ID pairs (e.g. device and sub-device)

Additionally, the `PCIID` struct is available for use to create one-off IDs.

Below is an example of querying for matching devices using a PCI ID pair:

```go
package main

import (
  "fmt"

  "github.com/powersj/pciids"
)

func main() {
  ids, err := pciids.QueryDevice("10de", "1467")
  if err != nil {
    fmt.Println("Error getting device info: %v", err)
  }

  for _, id := range ids {
    fmt.Println(id.String())
  }
}
```
