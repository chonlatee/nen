package cmd

import (
	"github.com/chonlatee/nen/cmd/crypto"
	"github.com/chonlatee/nen/cmd/gen"
	"github.com/chonlatee/nen/cmd/jwt"
	"github.com/spf13/cobra"
)

func init() {
	rootCMD.AddCommand(gen.GenCMD)
	rootCMD.AddCommand(crypto.CryptoCMD)
	rootCMD.AddCommand(jwt.JWTCMD)
}

var rootCMD = &cobra.Command{
	Use:   "nen [command]",
	Short: "A utility command line for developer",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func Excute() error {
	if err := rootCMD.Execute(); err != nil {
		return err
	}

	return nil
}
