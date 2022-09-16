package main

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"jwt-cli/pkg/token"
)

var _ Interface = (*SignOptions)(nil)
var opts = SignOptions{}

type SignOptions struct {
	PrivateKeyFile string
}

func Sign() *cobra.Command {
	o := &SignOptions{}

	cmd := &cobra.Command{
		Use:     "sign",
		Short:   "create signed JWT",
		Example: `jwt sign --key /path/to/key.pem <json to sign>`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key, err := token.LoadPem(opts.PrivateKeyFile)
			if err != nil {
				return err
			}

			content := []byte(args[0])
			var claims map[string]interface{}
			err = json.Unmarshal(content, &claims)
			if err != nil {
				log.Error(err)
				return err
			}
			jwt, err := token.Sign(key, claims)
			if err != nil {
				log.Error(err)
				return err
			}
			fmt.Printf("%s", jwt)
			return nil
		},
	}
	o.AddFlags(cmd)
	return cmd
}

func Verify() *cobra.Command {
	o := &SignOptions{}

	cmd := &cobra.Command{
		Use:     "verify",
		Short:   "verify signed JWT",
		Example: `jwt verify --key /path/to/key.pem <jwtstring>`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key, err := token.LoadPem(opts.PrivateKeyFile)
			if err != nil {
				return err
			}

			content := []byte(args[0])
			claims, err := token.Verify(&key.PublicKey, string(content))
			if err != nil {
				log.Error(err)
				return err
			}
			fmt.Printf("%s", string(claims))
			return nil
		},
	}
	o.AddFlags(cmd)
	return cmd
}

func (o *SignOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&opts.PrivateKeyFile, "key", "", "Path to private key file")
	err := cmd.MarkFlagRequired("key")
	if err != nil {
		log.Errorf("error marking flag required: %v", err)
		return
	}
}
