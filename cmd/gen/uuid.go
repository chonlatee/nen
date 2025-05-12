package gen

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chonlatee/nen/internal/nenuuid"
	"github.com/spf13/cobra"
)

var uuidCMD = &cobra.Command{
	Use:   "uuid [length]",
	Short: "Generate UUID v4. default is 10",
	RunE: func(cmd *cobra.Command, args []string) error {
		number := 10
		if len(args) >= 1 {
			input := args[0]
			num, err := strconv.Atoi(input)
			if err != nil {
				return err
			}

			number = num
		}

		var result string
		for range number {
			r, err := nenuuid.GenerateUUIDv4()
			if err != nil {
				fmt.Fprintf(os.Stdout, result)
				return err
			}
			result += r + "\n"
		}

		fmt.Fprintf(os.Stdout, result)

		return nil
	},
}
