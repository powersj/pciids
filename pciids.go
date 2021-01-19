package pciids

import (
	"github.com/powersj/pciids/pkg/file"
	"github.com/powersj/pciids/pkg/ids"
	"github.com/powersj/pciids/pkg/query"
)

// PCIID type.
type PCIID = ids.PCIID

// All returns all PCI IDs.
var All = query.All

// LatestFile returns the latest PCI IDs database file.
var LatestFile = file.Latest

// Parse will parse a PCI IDs file.
var Parse = ids.Parse

// QueryDevice will query using a specific PCI ID device pair.
var QueryDevice = query.Device

// QuerySubDevice will query using a specific PCI ID device and subdevice pair.
var QuerySubDevice = query.SubDevice
