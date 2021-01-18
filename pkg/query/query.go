package query

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/powersj/pciids/pkg/file"
	"github.com/powersj/pciids/pkg/ids"
)

// All returns all PCI IDs in a slice.
func All() ([]ids.PCIID, error) {
	rawIDs, err := file.Latest()
	if err != nil {
		return nil, errors.Wrap(err, "Cannot read latest IDs")
	}

	parsedIDs, err := ids.Parse(rawIDs)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot parse IDs")
	}

	return parsedIDs, nil
}

// Device finds all matching devices based on a pair of vendor and device IDs.
func Device(vendorID string, deviceID string) ([]ids.PCIID, error) {
	var results []ids.PCIID = make([]ids.PCIID, 0)

	var vID string = strings.ToLower(vendorID)
	var dID string = strings.ToLower(deviceID)

	pciids, err := All()
	if err != nil {
		return results, err
	}

	// check if the device first has sub IDs. If so see if the pair we are
	// looking for are contained there. If no sub IDs, then check the main IDs.
	for _, id := range pciids {
		if id.SubVendorID != "" && id.SubDeviceID != "" {
			if vID == id.SubVendorID && dID == id.SubDeviceID {
				results = append(results, id)
			}
		} else if vID == id.VendorID && dID == id.DeviceID {
			results = append(results, id)
		}
	}

	return results, nil
}

// SubDevice finds all matching devices based on a quartet of IDs.
func SubDevice(
	vendorID string, deviceID string, subVendorID string, subDeviceID string) (
	[]ids.PCIID, error) {
	var results []ids.PCIID = make([]ids.PCIID, 0)

	var vID string = strings.ToLower(vendorID)
	var dID string = strings.ToLower(deviceID)
	var sVID string = strings.ToLower(subVendorID)
	var sDID string = strings.ToLower(subDeviceID)

	pciids, err := All()
	if err != nil {
		return results, err
	}

	for _, id := range pciids {
		if vID == id.VendorID && dID == id.DeviceID {
			if sVID == id.SubVendorID && sDID == id.SubDeviceID {
				results = append(results, id)
			}
		}
	}

	return results, nil
}
