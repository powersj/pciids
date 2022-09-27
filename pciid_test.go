package pciids

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPCIID(t *testing.T) {
	t.Parallel()

	empty := PCIID{}
	pciid := PCIID{
		VendorID:   "5678",
		VendorName: "Kraftworks",
		DeviceID:   "1234",
		DeviceName: "Foobar 9000",
	}

	assert.Contains(t, pciid.String(), "Kraftworks Foobar 9000")
	assert.NoError(t, json.Unmarshal([]byte(pciid.JSON()), &empty))
}

func TestSubPCIID(t *testing.T) {
	t.Parallel()

	empty := PCIID{}
	pciid := PCIID{
		VendorID:   "5678",
		VendorName: "Kraftworks",
		DeviceID:   "1234",
		DeviceName: "Foobar 9000",

		SubVendorID:   "9000",
		SubVendorName: "Vendor",
		SubDeviceID:   "0000",
		SubDeviceName: "Company",
	}

	assert.Contains(t, pciid.String(), "Vendor Company")
	assert.NoError(t, json.Unmarshal([]byte(pciid.JSON()), &empty))
}
