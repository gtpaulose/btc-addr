package cmd

import (
	"fmt"

	"github.com/gtpaulose/btc-addr/internal/multisig"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// multisigCmd represents the multisig command
var multisigCmd = &cobra.Command{
	Use:   "multisig [OPTIONS] public_keys...",
	Short: "Generates a n-of-m multisig address",
	Long:  `By providing the values of n, m and the public keys, the client will return a multi-sig address which can be safely distributed`,
	Run: func(cmd *cobra.Command, args []string) {
		pubkeys := [][]byte{}
		for _, key := range args {
			pubkeys = append(pubkeys, []byte(key))
		}

		address, err := multisig.BuildP2SHAddress(cast.ToInt(cmd.Flag("n").Value.String()), cast.ToInt(cmd.Flag("m").Value.String()), pubkeys)
		if err != nil {
			fmt.Println("Error building multi sig address: ", err)
			return
		}

		fmt.Println("\nAddress: ", address)
	},
}

func init() {
	multisigCmd.Flags().Int("n", 0, "min signers")
	multisigCmd.MarkFlagRequired("n")
	multisigCmd.Flags().Int("m", 0, "total signers")
	multisigCmd.MarkFlagRequired("m")
}
