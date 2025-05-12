package jwt

import (
	"errors"
	"fmt"
	"os"

	"github.com/chonlatee/nen/internal/nenjwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

func init() {
	JWTCMD.AddCommand(signJwt)
	JWTCMD.AddCommand(decodeJwt)
	JWTCMD.AddCommand(verifyJwt)
	signJwt.Flags().StringP("payload", "", "", "path payload json file format")
	signJwt.Flags().StringP("key", "", "", "path to private key file pem format")
	signJwt.MarkFlagRequired("payload")
	signJwt.MarkFlagRequired("key")
	verifyJwt.Flags().StringP("key", "", "", "path to public key file pem format")
	verifyJwt.MarkFlagRequired("key")
}

var JWTCMD = &cobra.Command{
	Use:   "jwt",
	Short: "A jwt utility command line tool for sign, verify and decode",
}

var signJwt = &cobra.Command{
	Use:     "sign",
	Short:   "Generate jwt token from payload file and private key file.",
	Example: "nen jwt sign --payload=payload.json --key=private.pem",
	RunE: func(cmd *cobra.Command, args []string) error {
		payloadInput, err := cmd.Flags().GetString("payload")
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return err
		}

		if _, err := os.Stat(payloadInput); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "payload file not exist\n")
			return err
		}

		p, err := os.ReadFile(payloadInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read payload file err: %v\n", err)
			return err
		}

		keyFileInput, err := cmd.Flags().GetString("key")
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return err
		}

		if _, err := os.Stat(keyFileInput); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "private key file not exist\n")
			return err
		}

		k, err := os.ReadFile(keyFileInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read private key file err: %v\n", err)
			return err
		}

		r, err := nenjwt.Sign(p, k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "sign jwt error: %v\n", err)
			return err
		}

		fmt.Fprintf(os.Stdout, r+"\n")

		return nil

	},
}

var decodeJwt = &cobra.Command{
	Use:     "decode",
	Short:   "Decode jwt header and payload without verify",
	Example: "nen jwt decode x.y.z",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) <= 0 {
			return cmd.Help()
		}

		input := args[0]

		header, body, err := nenjwt.Decode(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "decode jwt error: %v\n", err.Error())
			return err
		}

		fmt.Fprintf(os.Stdout, header+"\n")
		fmt.Fprintf(os.Stdout, body+"\n")

		return nil
	},
}

var verifyJwt = &cobra.Command{
	Use:     "verify",
	Short:   "Verify jwt with public key file",
	Example: "nen jwt verify x.y.z --key=public.pem",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) <= 0 {
			return cmd.Help()
		}

		token := args[0]

		keyInput, err := cmd.Flags().GetString("key")
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return err
		}

		if _, err := os.Stat(keyInput); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "public key file not found")
			return err
		}

		key, err := os.ReadFile(keyInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read public key error: %v\n", err)
			return err
		}

		jwtToken, err := nenjwt.Verify(token, key)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				fmt.Fprintf(os.Stderr, "jwt token expired\n")
				return nil
			}
			fmt.Fprintf(os.Stderr, "verify jwt error: %v\n", err)
			return err
		}

		if !jwtToken.Valid {
			fmt.Fprintf(os.Stdout, "jwt token not valid \n")
			return nil
		}

		fmt.Fprintf(os.Stdout, "jwt token valid\n")

		return nil

	},
}
