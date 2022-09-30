package pciids

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	remoteURL       = "https://raw.githubusercontent.com/pciutils/pciids/master/pci.ids"
	localPathEnvVar = "PCIIDS_LOCAL_PATH"
)

// All returns all PCI IDs in a slice.
func All() ([]PCIID, error) {
	var rawIDs string
	var err error

	path, exist := os.LookupEnv(localPathEnvVar)
	if exist {
		rawIDs, err = Local(path)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read local IDs")
		}
	} else {
		rawIDs, err = Latest()
		if err != nil {
			return nil, errors.Wrap(err, "cannot read latest IDs")
		}
	}

	parsedIDs := parse(rawIDs)

	return parsedIDs, nil
}

func Local(path string) (string, error) {
	log.Debug("reading: ", path)

	if _, err := os.Stat(path); err != nil {
		return "", errors.Wrap(err, "could not stat local file")
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", errors.Wrap(err, "could not read local file")
	}

	return string(bytes), nil
}

// Latest downloads the latest PCI ID file from the GitHub mirror.
func Latest() (string, error) {
	return LatestWithContext(context.Background())
}

func LatestWithContext(ctx context.Context) (string, error) {
	log.Debug("downloading ", remoteURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, remoteURL, nil)
	if err != nil {
		return "", errors.Wrap(err, "http request error")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "http response error")
	}
	defer resp.Body.Close()

	log.Debug(resp.Status)

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintln("invalid response status code: ", resp.Status))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "could not read downloaded file")
	}

	return string(bodyBytes), nil
}

func QueryVendor(vendorID string) ([]PCIID, error) {
	results := make([]PCIID, 0)

	vID := strings.ToLower(vendorID)

	log.Debug(fmt.Sprintf("looking up %s", vID))

	pciids, err := All()
	if err != nil {
		return results, err
	}

	for _, pciid := range pciids {
		if vID == pciid.VendorID {
			results = append(results, pciid)
		}
	}

	log.Debug(fmt.Sprintf("found %d result(s)", len(results)))

	return results, nil
}

// Device finds all matching devices based on a pair of vendor and device IDs.
func QueryDevice(vendorID string, deviceID string) ([]PCIID, error) {
	results := make([]PCIID, 0)

	vID := strings.ToLower(vendorID)
	dID := strings.ToLower(deviceID)

	log.Debug(fmt.Sprintf("looking up %s:%s", vID, dID))

	pciids, err := All()
	if err != nil {
		return results, err
	}

	// check if the device first has sub IDs. If so see if the pair we are
	// looking for are contained there. If no sub IDs, then check the main IDs.
	for _, pciid := range pciids {
		if pciid.SubVendorID != "" && pciid.SubDeviceID != "" {
			if vID == pciid.SubVendorID && dID == pciid.SubDeviceID {
				results = append(results, pciid)
			}
		} else if vID == pciid.VendorID && dID == pciid.DeviceID {
			results = append(results, pciid)
		}
	}

	log.Debug(fmt.Sprintf("found %d result(s)", len(results)))

	return results, nil
}

// SubDevice finds all matching devices based on a quartet of IDs.
func QuerySubDevice(
	vendorID string, deviceID string, subVendorID string, subDeviceID string) (
	[]PCIID, error,
) {
	results := make([]PCIID, 0)

	vID := strings.ToLower(vendorID)
	dID := strings.ToLower(deviceID)
	sVID := strings.ToLower(subVendorID)
	sDID := strings.ToLower(subDeviceID)

	log.Debug(fmt.Sprintf("looking up %s:%s %s:%s", vID, dID, sVID, sDID))

	pciids, err := All()
	if err != nil {
		return results, err
	}

	for _, pciid := range pciids {
		if vID == pciid.VendorID && dID == pciid.DeviceID {
			if sVID == pciid.SubVendorID && sDID == pciid.SubDeviceID {
				results = append(results, pciid)
			}
		}
	}

	log.Debug(fmt.Sprintf("found %d result(s)", len(results)))

	return results, nil
}
