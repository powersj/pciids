package main

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// PCIID struct.
type PCIID struct {
	VendorID   string `json:"vendorId"`
	VendorName string `json:"vendorName"`
	DeviceID   string `json:"deviceId"`
	DeviceName string `json:"deviceName"`

	SubVendorID   string `json:"subvendorId,omitempty"`
	SubVendorName string `json:"subvendorName,omitempty"`
	SubDeviceID   string `json:"subdeviceId,omitempty"`
	SubDeviceName string `json:"subdeviceName,omitempty"`
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
		"%s:%s - %s %s", p.VendorID, p.DeviceID, p.VendorName, p.DeviceName,
	)
}

// JSON representation of the struct.
func (p *PCIID) JSON() string {
	bytes, err := json.MarshalIndent(&p, "", "  ")
	if err != nil {
		log.Error("error marshaling JSON:", err)

		return ""
	}

	return string(bytes)
}
