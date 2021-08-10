package cmd

import (
	"fmt"

	"github.com/gtpaulose/btc-addr/pkg/keystore"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List addresses [OPTIONS]",
	Long: `List addresses for different formats by providing a mnemonic or a root key. 
The mnemonic must be enclosed in double quotes and the root key should be in the BIP32 format. 
	`,
	Example: `  btc-addr list -m "mask capable giant patient subject buffalo lab armed potato twice barrel online"
  btc-addr list -m "creek multiply disorder edge hotel armor labor kidney remain outdoor orange spoon" -f bip49 -p "m/49'/0'/0'/0/1"`,
	Run: func(cmd *cobra.Command, args []string) {
		mnemonic := cmd.Flag("mnemonic").Value.String()
		pp := cmd.Flag("passphrase").Value.String()
		path := cmd.Flag("path").Value.String()
		num := cast.ToInt(cmd.Flag("num").Value.String())

		purpose, err := keystore.ValidateFormat(cmd.Flag("format").Value.String())
		if err != nil {
			fmt.Printf("%s\n\n", err.Error())
			cmd.Help()
			return
		}

		ks := keystore.New(keystore.WithMnemonic(mnemonic), keystore.WithPassphrase(pp))
		if path == "" {
			ks.PrintAddresses(purpose, num)
			return
		}

		if err := keystore.ValidatePath(path, purpose); err != nil {
			fmt.Printf("%s\n\n", err.Error())
			cmd.Help()
			return
		}
		ks.PrintAddressForPath(purpose, path)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("format", "f", "bip44", "format of the address (eg. bip44, bip49, bip84)")
	listCmd.Flags().IntP("num", "n", 1, "number of addresses to list starting from path")
	listCmd.Flags().StringP("mnemonic", "m", "", "mnemonic of account")
	listCmd.MarkFlagRequired("mnemonic")
	listCmd.Flags().String("passphrase", "", "passphrase associated with the mnemonic")
	listCmd.Flags().StringP("path", "p", "", "reference path to list addresses")
}
