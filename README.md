# pciids

*Lookup vendor and device names using PCI IDs!*

[![Build Status](https://travis-ci.com/powersj/pciids.svg?branch=master)](https://travis-ci.com/powersj/pciids/) [![Go Report Card](https://goreportcard.com/badge/github.com/powersj/pciids)](https://goreportcard.com/report/github.com/powersj/pciids) [![Go Reference](https://pkg.go.dev/badge/github.com/powersj/pciids.svg)](https://pkg.go.dev/github.com/powersj/pciids)

## CLI

To search for devices using the CLI, pass in either:

  a) a pair of vendor and device PCI IDs
  b) two pairs, vendor and device PCI IDs as well as sub-vendor and
     sub-device PCI IDs:

Here are some examples:

```text
$ pciids 1ed5
1ed5:0100 - Moore Threads Technology Co.,Ltd MTT S10
1ed5:0101 - Moore Threads Technology Co.,Ltd MTT S10
1ed5:0102 - Moore Threads Technology Co.,Ltd MTT S30
1ed5:0105 - Moore Threads Technology Co.,Ltd MTT S50
1ed5:0106 - Moore Threads Technology Co.,Ltd MTT S60
1ed5:0111 - Moore Threads Technology Co.,Ltd MTT S100
1ed5:0121 - Moore Threads Technology Co.,Ltd MTT S1000M
1ed5:0122 - Moore Threads Technology Co.,Ltd MTT S1000
1ed5:0123 - Moore Threads Technology Co.,Ltd MTT S2000
1ed5:01ff - Moore Threads Technology Co.,Ltd MTT HDMI/DP Audio
1ed5:0201 - Moore Threads Technology Co.,Ltd G2D30
1ed5:0202 - Moore Threads Technology Co.,Ltd G2D20
1ed5:0203 - Moore Threads Technology Co.,Ltd G2D10
1ed5:0211 - Moore Threads Technology Co.,Ltd G2D40
1ed5:0221 - Moore Threads Technology Co.,Ltd G2S80
1ed5:0222 - Moore Threads Technology Co.,Ltd G2S85
1ed5:0223 - Moore Threads Technology Co.,Ltd G2S4
1ed5:0251 - Moore Threads Technology Co.,Ltd G2N10
1ed5:02ff - Moore Threads Technology Co.,Ltd MTT HDMI/DP Audio
$ pciids 1d0f efa1
1d0f:efa1 - Amazon.com, Inc. Elastic Fabric Adapter (EFA)
$ pciids 10de 2206 10de 1467
10de:2206 10de:1467 - NVIDIA Corporation GA102 [GeForce RTX 3080]
```

If there are multiple matches then all matches are returned.

### JSON output

The command can take the `--json` flag to produce JSON output:

```shell
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

```shell
$ pciids --json 121a 0009 121a 0009
DEBU looking up 121a:0009 121a:0009
DEBU downloading https://raw.githubusercontent.com/pciutils/pciids/master/pci.ids
DEBU 200 OK
DEBU parsing vendor IDs
DEBU parsing PCI IDs
DEBU found 1 results
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
* `QueryVendor(vendorID)`: Searches for devices matching a vendor ID
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
