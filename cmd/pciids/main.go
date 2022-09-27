package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/powersj/pciids/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	version         = "v2.3.0"
	numVendorIDs    = 1
	numDeviceIDs    = 2
	numSubDeviceIDs = 4
)

var (
	debugOutput bool
	jsonOutput  bool
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "pciids",
	Short: "Lookup vendor and device names using PCI IDs",
	Long: `Lookup vendor and device names using PCI IDs.

To search for devices using the CLI, pass in either:
  a) vendor
  b) a pair of vendor and device PCI IDs
  c) two pairs, vendor and device PCI IDs as well as sub-vendor and
     sub-device PCI IDs:

Examples:
$ pciids 1ed5
$ pciids 1d0f efa1
$ pciids 10de 2206 10de 1467
`,
	PersistentPreRun: setup,
	Args:             args,
	RunE:             root,
}

// Called before all commands to setup general run-time settings.
func setup(cmd *cobra.Command, args []string) {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	})

	if debugOutput {
		log.SetLevel(log.DebugLevel)
	}
}

// Checks that there are exactly 1, 2, or 4 arguments.
func args(cmd *cobra.Command, args []string) error {
	length := len(args)
	if length != numVendorIDs && length != numDeviceIDs &&
		length != numSubDeviceIDs {
		return errors.New(
			"invalid number of arguments, expected 1, 2, or 4 (e.g. 10de 1467)",
		)
	}

	return nil
}

// Base command operations.
func root(cmd *cobra.Command, args []string) error {
	var ids []pciids.PCIID
	var err error

	switch len(args) {
	case numSubDeviceIDs:
		ids, err = pciids.QuerySubDevice(args[0], args[1], args[2], args[3])
	case numDeviceIDs:
		ids, err = pciids.QueryDevice(args[0], args[1])
	case numVendorIDs:
		ids, err = pciids.QueryVendor(args[0])
	}

	if err != nil {
		return errors.Wrap(err, "Error while querying for device")
	}

	if jsonOutput {
		b, err := json.MarshalIndent(&ids, "", "  ")
		if err != nil {
			return errors.Wrap(err, "Unable to convert to JSON")
		}

		fmt.Println(string(b))
	} else {
		for _, id := range ids {
			fmt.Println(id.String())
		}
	}

	return nil
}

// CLI function to setup flags.
func init() {
	rootCmd.Version = version

	rootCmd.PersistentFlags().BoolVar(
		&debugOutput, "debug", false, "debug output",
	)
	rootCmd.PersistentFlags().BoolVar(
		&jsonOutput, "json", false, "enable output in JSON",
	)
}

// Execute adds all child commands to the root command and sets flags.
//
// This is called by main.main(). It only needs to happen once to the rootCmd.
func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
