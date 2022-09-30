package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/powersj/pciids/v2"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	version         = "v2.3.0"
	numVendorIDs    = 1
	numDeviceIDs    = 2
	numSubDeviceIDs = 4
	usage           = "Lookup vendor and device names using PCI IDs"
	usageText       = "pciids [global options] vendorID [deviceID [subdeviceID subvendorID]]"
	description     = `To search for devices using the CLI, pass in either:

* vendor ID
* a pair of vendor and device PCI IDs
* two pairs, vendor and device PCI IDs as well as sub-vendor and sub-device PCI IDs

Examples:
$ pciids 1ed5
$ pciids 1d0f efa1
$ pciids 10de 2206 10de 1467
`
)

var (
	debugOutput bool
	jsonOutput  bool
)

// Called before all commands to setup general run-time settings.
func setupLogging() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	})

	if debugOutput {
		log.SetLevel(log.DebugLevel)
	}
}

// Checks that there are exactly 1, 2, or 4 arguments.
func verifyArgs(args []string) error {
	length := len(args)
	if length != numVendorIDs && length != numDeviceIDs &&
		length != numSubDeviceIDs {
		return errors.New(
			"invalid number of arguments, expected 1, 2, or 4 (e.g. 10de 1467)",
		)
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags.
//
// This is called by main.main(). It only needs to happen once to the rootCmd.
func main() {
	app := &cli.App{
		Name:        "pciids",
		Version:     version,
		Usage:       usage,
		UsageText:   usageText,
		Description: description,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "debug output",
				Destination: &debugOutput,
			},
			&cli.BoolFlag{
				Name:        "json",
				Usage:       "enable output in JSON",
				Destination: &jsonOutput,
			},
		},
		Action: func(cCtx *cli.Context) error {
			setupLogging()

			args := cCtx.Args().Slice()
			err := verifyArgs(args)
			if err != nil {
				return err
			}

			var ids []pciids.PCIID
			switch len(args) {
			case numSubDeviceIDs:
				ids, err = pciids.QuerySubDevice(args[0], args[1], args[2], args[3])
			case numDeviceIDs:
				ids, err = pciids.QueryDevice(args[0], args[1])
			case numVendorIDs:
				ids, err = pciids.QueryVendor(args[0])
			}
			if err != nil {
				return errors.Wrap(err, "error while querying for device")
			}

			if jsonOutput {
				b, err := json.MarshalIndent(&ids, "", "  ")
				if err != nil {
					return errors.Wrap(err, "unable to convert to JSON")
				}

				fmt.Println(string(b))
			} else {
				for _, id := range ids {
					fmt.Println(id.String())
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
