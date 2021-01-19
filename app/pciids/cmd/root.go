package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/powersj/pciids/pkg/ids"
	"github.com/powersj/pciids/pkg/query"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	version         = "v1.0.0"
	numDeviceIDs    = 2
	numSubDeviceIDs = 4
)

var (
	debugOutput bool
	jsonOutput  bool
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:              "pciids",
	Short:            "Lookup vendor and device names using PCI IDs",
	Long:             `Lookup vendor and device names using PCI IDs`,
	PersistentPreRun: setup,
	Args:             args,
	RunE:             root,
}

// Called before all commands to setup general run-time settings.
func setup(cmd *cobra.Command, args []string) {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
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
				"require either two or four PCI IDs as arguments (e.g. 10de 1467)",
			)
		}
	}

	return nil
}

// Base command operations.
func root(cmd *cobra.Command, args []string) error {
	var ids []ids.PCIID
	var err error

	if len(args) == numSubDeviceIDs {
		ids, err = query.SubDevice(args[0], args[1], args[2], args[3])
		if err != nil {
			return errors.Wrap(err, "Error while querying for device")
		}
	} else {
		ids, err = query.Device(args[0], args[1])
		if err != nil {
			return errors.Wrap(err, "Error while querying for device")
		}
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
