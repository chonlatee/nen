package gen

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chonlatee/nen/internal/nenrand"
	"github.com/spf13/cobra"
)

func init() {
	randStrCMD.AddCommand(randHexCMD)
}

var randStrCMD = &cobra.Command{
	Use:   "randstr",
	Short: "Generate random string",
}

var randHexCMD = &cobra.Command{
	Use:   "hex [length]",
	Short: "Generate random hex string from length",
	RunE: func(cmd *cobra.Command, args []string) error {
		length := 32
		if len(args) >= 1 {
			input := args[0]
			num, err := strconv.Atoi(input)
			if err != nil {
				return err
			}
			length = num
		}

		r, err := nenrand.GenerateHex(length)
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", r)

		return nil
	},
}
