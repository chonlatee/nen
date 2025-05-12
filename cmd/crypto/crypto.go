package crypto

import (
	"fmt"
	"os"

	"github.com/chonlatee/nen/internal/nencrypto"
	"github.com/spf13/cobra"
)

func init() {
	CryptoCMD.AddCommand(rsaKeyCMD)
	rsaKeyCMD.Flags().BoolP("file", "", false, "output to file")
}

var CryptoCMD = &cobra.Command{
	Use:   "crypto",
	Short: "A utility for cryptography  eg. generate RSA key ",
}

var rsaKeyCMD = &cobra.Command{
	Use:     "rsa",
	Short:   "Generate RSA key",
	Example: "nen crypto rsa, nen crypto rsa --file",
	RunE: func(cmd *cobra.Command, args []string) error {

		pk, pb, err := nencrypto.GenerateRSAKeyPEM()
		if err != nil {
			return err
		}

		out, err := cmd.Flags().GetBool("file")
		if out {
			err := os.WriteFile("private.pem", []byte(pk), 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "create private key file error: %v", err)
				return err
			}

			err = os.WriteFile("public.pem", []byte(pb), 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "create public key file error: %v", err)
				return err
			}

			fmt.Fprintf(os.Stdout, "create private key and public key file successfully\n")

		} else {
			fmt.Fprintf(os.Stdout, "%s\n%s", pk, pb)
		}

		return nil
	},
}
