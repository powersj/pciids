package pciids

import (
	"bufio"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	currentDevice   PCIID
	currentVendorID string
	vendors         map[string]string
)

// Parse returns a list of all PCI IDs in a slice after parsing the PCI ID file.
func parse(rawIDs string) []PCIID {
	var devices []PCIID

	if vendors == nil {
		vendors = parseVendors(rawIDs)
	}

	log.Debug("parsing PCI IDs")

	scanner := bufio.NewScanner(strings.NewReader(rawIDs))
	for scanner.Scan() {
		line := scanner.Text()
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
			splits := strings.Split(line, "  ")
			currentVendorID = strings.TrimSpace(splits[0])
		}
	}

	return devices
}

// parseSubDevice parses a line starting with two tabs and returns the PCI ID.
func parseSubDevice(line string) PCIID {
	splits := strings.Split(strings.TrimPrefix(line, "\t\t"), " ")
	subVendorID := strings.TrimSpace(splits[0])

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
	splits := strings.Split(strings.TrimPrefix(line, "\t"), "  ")

	return PCIID{
		DeviceID:   strings.TrimSpace(splits[0]),
		DeviceName: strings.TrimSpace(strings.Join(splits[1:], " ")),
		VendorID:   currentVendorID,
		VendorName: vendors[currentVendorID],
	}
}

// parseVendors parses all Vendors from the PCI ID file.
func parseVendors(rawIDs string) map[string]string {
	vendors := make(map[string]string)

	log.Debug("parsing vendor IDs")

	scanner := bufio.NewScanner(strings.NewReader(rawIDs))
	for scanner.Scan() {
		line := scanner.Text()
		if !validVendor(line) {
			continue
		}

		splits := strings.Split(line, "  ")
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
