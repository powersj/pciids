package ids

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPCIID(t *testing.T) {
	var id PCIID = PCIID{
		VendorID:   "5678",
		VendorName: "Kraftworks",
		DeviceID:   "1234",
		DeviceName: "Foobar 9000",
	}

	assert.Contains(t, id.String(), "Kraftworks Foobar 9000")

	var newID PCIID
	assert.NoError(t, json.Unmarshal([]byte(id.JSON()), &newID))
}

func TestSubPCIID(t *testing.T) {
	var id PCIID = PCIID{
		VendorID:   "5678",
		VendorName: "Kraftworks",
		DeviceID:   "1234",
		DeviceName: "Foobar 9000",

		SubVendorID:   "9000",
		SubVendorName: "Vendor",
		SubDeviceID:   "0000",
		SubDeviceName: "Company",
	}

	assert.Contains(t, id.String(), "Vendor Company")

	var newID PCIID
	assert.NoError(t, json.Unmarshal([]byte(id.JSON()), &newID))
}
