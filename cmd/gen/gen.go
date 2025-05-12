package gen

import "github.com/spf13/cobra"

func init() {
	GenCMD.AddCommand(uuidCMD)
	GenCMD.AddCommand(randStrCMD)
}

var GenCMD = &cobra.Command{
	Use:   "gen",
	Short: "A utility for generate something.",
}
