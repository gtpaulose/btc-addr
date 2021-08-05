package cmd

import (
	"github.com/gtpaulose/btc-addr/internal/keystore"
	"github.com/spf13/cobra"
)

// mnemonicCmd represents the mnemonic command
var mnemonicCmd = &cobra.Command{
	Use:   "mnemonic",
	Short: "Generate bip39 mnemonics and associated keys",
	Run: func(cmd *cobra.Command, args []string) {
		keystore.New(keystore.GenerateMnemonic(12), keystore.WithPassphrase(cmd.Flag("passphrase").Value.String())).Print()
	},
}

func init() {
	mnemonicCmd.Flags().StringP("passphrase", "p", "", "passphrase")
}
