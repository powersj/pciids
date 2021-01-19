package file

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	remoteURL = "https://raw.githubusercontent.com/pciutils/pciids/master/pci.ids"
)

// Testing is used to return the test string, rather than external value.
var Testing bool = false

var testString string = "121a  3Dfx Interactive, Inc.\n" +
	"\t0001  Voodoo\n" +
	"\t0002  Voodoo 2\n" +
	"\t0003  Voodoo Banshee\n" +
	"\t\t121a 0001  Voodoo Banshee AGP\n" +
	"\t\t121a 0003  Voodoo Banshee AGP SGRAM\n" +
	"\t0009  Voodoo 4 / Voodoo 5\n" +
	"\t\t121a 0003  Voodoo5 PCI 5500\n" +
	"\t\t121a 0009  Voodoo5 AGP 5500/6000\n"

// Latest downloads the latest PCI ID file from the GitHub mirror.
func Latest() (string, error) {
	if Testing {
		log.Debug("Using test string")
		return testString, nil
	}

	log.Debug("Downloading ", remoteURL)
	resp, err := http.Get(remoteURL)
	if err != nil {
		return "", errors.Wrap(err, "Failed to download latest file")
	}
	defer resp.Body.Close()

	log.Debug(resp.Status)
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintln("Invalid response status code: ", resp.Status))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Could not read downloaded file")
	}

	return string(bodyBytes), nil
}
