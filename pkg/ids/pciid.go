package ids

import (
	"encoding/json"
	"fmt"
)

// PCIID struct.
type PCIID struct {
	VendorID   string `json:"vendorID"`
	DeviceID   string `json:"deviceID"`
	VendorName string `json:"vendorName"`
	DeviceName string `json:"deviceName"`

	SubVendorID   string `json:"subVendorID"`
	SubDeviceID   string `json:"subDeviceID"`
	SubVendorName string `json:"subVendorName"`
	SubDeviceName string `json:"subDeviceName"`
}

// String representation of the struct.
func (p *PCIID) String() string {
	if (p.SubVendorID != "") && (p.SubDeviceID != "") {
		return fmt.Sprintf(
			"%s:%s %s:%s - %s %s",
			p.VendorID, p.DeviceID, p.SubVendorID, p.SubDeviceID, p.SubVendorName, p.SubDeviceName,
		)
	}

	return fmt.Sprintf(
		"%s:%s           - %s %s", p.VendorID, p.DeviceID, p.VendorName, p.DeviceName,
	)
}

// JSON representation of the struct.
func (p *PCIID) JSON() string {
	return objectJSONString(&p)
}

// objectJSONString converts a struct to a JSON object with 2 space indent.
func objectJSONString(o interface{}) string {
	b, err := json.MarshalIndent(&o, "", "  ")
	if err != nil {
		fmt.Println("error marshaling JSON:", err)
		return ""
	}

	return string(b)
}
