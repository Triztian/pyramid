package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// Generate pyramid command
var (
	// Flag pointers
	minRand int64
	maxRand int64

	genCmd = &cobra.Command{
		Use:   "generate [levels]",
		Short: "Generate a pyramid and print it",
		Long: `
A pyramid is represented as such:
	1
	2 3
	4 5 6

Examples:
	$ pyramid generate 10 
	$ pyramid generate --min 10 --max 20 10
		`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("missing levels argument")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			lvs, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return nil
			}

			ns := generate(int(lvs), int(minRand), int(maxRand))

			return writeTo(os.Stdout, ns)
		},
	}
)

func init() {
	genCmd.Flags().Int64Var(&minRand, "min", DefaultMin, "The minimum possible value for a node")

	genCmd.Flags().Int64Var(&maxRand, "max", DefaultMax, "The maximum possible value for a node")
}
