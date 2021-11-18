package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
  a) a pair of vendor and device PCI IDs
  b) two pairs, vendor and device PCI IDs as well as sub-vendor and
     sub-device PCI IDs:

Examples:
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

// Checks that there are exactly two arguments.
func args(cmd *cobra.Command, args []string) error {
	if len(args) != numDeviceIDs {
		if len(args) != numSubDeviceIDs {
			return errors.New(
				"either two or four PCI IDs as arguments (e.g. 10de 1467) are required",
			)
		}
	}

	return nil
}

// Base command operations.
func root(cmd *cobra.Command, args []string) error {
	var ids []PCIID
	var err error

	if len(args) == numSubDeviceIDs {
		ids, err = QuerySubDevice(args[0], args[1], args[2], args[3])
	} else {
		ids, err = QueryDevice(args[0], args[1])
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
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
