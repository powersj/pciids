package ids

import (
	"bufio"
	"strings"
)

var (
	currentDevice   PCIID
	currentVendorID string
	vendors         map[string]string
)

// Parse returns a list of all PCI IDs in a slice after parsing the PCI ID file.
func Parse(rawIDs string) ([]PCIID, error) {
	var devices []PCIID

	if vendors == nil {
		vendors = parseVendors(rawIDs)
	}

	scanner := bufio.NewScanner(strings.NewReader(rawIDs))
	for scanner.Scan() {
		var line string = scanner.Text()
		if !valid(line) {
			continue
		}

		switch {
		case strings.HasPrefix(line, "\t\t"):
			devices = append(devices, parseSubDevice(line))
		case strings.HasPrefix(line, "\t"):
			currentDevice = parseDevice(line)
			devices = append(devices, currentDevice)
		default:
			var splits []string = strings.Split(line, "  ")
			currentVendorID = strings.TrimSpace(splits[0])
		}
	}

	return devices, nil
}

// parseSubDevice parses a line starting with two tabs and returns the PCI ID.
func parseSubDevice(line string) PCIID {
	var splits []string = strings.Split(strings.TrimPrefix(line, "\t\t"), " ")
	var subVendorID string = strings.TrimSpace(splits[0])

	return PCIID{
		DeviceID:      currentDevice.DeviceID,
		DeviceName:    currentDevice.DeviceName,
		VendorID:      currentDevice.VendorID,
		VendorName:    currentDevice.VendorName,
		SubDeviceID:   splits[1],
		SubDeviceName: strings.Join(splits[3:], " "),
		SubVendorID:   subVendorID,
		SubVendorName: vendors[subVendorID],
	}
}

// parseDevice parses the line starting with a tab and returns the PCI ID.
func parseDevice(line string) PCIID {
	var splits []string = strings.Split(strings.TrimPrefix(line, "\t"), "  ")

	return PCIID{
		DeviceID:   strings.TrimSpace(splits[0]),
		DeviceName: strings.TrimSpace(strings.Join(splits[1:], " ")),
		VendorID:   currentVendorID,
		VendorName: vendors[currentVendorID],
	}
}

// parseVendors parses all Vendors from the PCI ID file.
func parseVendors(rawIDs string) map[string]string {
	var vendors map[string]string = make(map[string]string)

	scanner := bufio.NewScanner(strings.NewReader(rawIDs))
	for scanner.Scan() {
		var line string = scanner.Text()
		if !validVendor(line) {
			continue
		}

		var splits []string = strings.Split(line, "  ")
		vendors[strings.TrimSpace(splits[0])] = strings.Join(splits[1:], " ")
	}

	return vendors
}

// valid returns boolean if the line is a valid PCI ID content.
func valid(line string) bool {
	switch {
	case len(line) == 0:
		return false
	case strings.HasPrefix(line, "#"):
		return false
	case strings.HasPrefix(line, "C"):
		return false
	}

	return true
}

// validVendor returns boolean if the line is a valid PCI ID vendor line.
func validVendor(line string) bool {
	switch {
	case len(line) == 0:
		return false
	case strings.HasPrefix(line, "\t"):
		return false
	case strings.HasPrefix(line, "#"):
		return false
	case strings.HasPrefix(line, "C"):
		return false
	}

	return true
}
