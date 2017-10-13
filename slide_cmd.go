package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var slideCmd = &cobra.Command{
	Use:   "slide [input.txt]",
	Short: "Calulates the quickest path down the pyramid.",
	Long: `
Examples:

	$ pyramid slide /path/to/challenge.txt
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing file")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}

		pyramid, err := scanReader(f)
		if err != nil {
			return err
		}

		ans := shortest(pyramid[1:])
		levelCount := pyramid[0][0]

		fmt.Printf("Answer: %d -> %d\n", levelCount, ans)

		return nil
	},
}
