package cmd

import (
	"github.com/gtpaulose/btc-addr/internal/keystore"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// mnemonicCmd represents the mnemonic command
var mnemonicCmd = &cobra.Command{
	Use:   "mnemonic",
	Short: "Generate bip39 mnemonics and associated keys",
	Long:  "Generate bip39 mnemonics for a given phrase length. Recommended values for phrases are 12,15,18,21 and 24",
	Run: func(cmd *cobra.Command, args []string) {
		keystore.New(keystore.GenerateMnemonic(cast.ToInt(cmd.Flag("num").Value.String())), keystore.WithPassphrase(cmd.Flag("passphrase").Value.String())).Print()
	},
}

func init() {
	mnemonicCmd.Flags().StringP("passphrase", "p", "", "passphrase")
	mnemonicCmd.Flags().Int("num", 12, "number of phrases in mnemonic")
}
