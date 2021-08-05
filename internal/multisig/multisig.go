package multisig

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

// BuildP2SHAddress creates a N-of-M multisig address given m, n and n public keys
func BuildP2SHAddress(n, m int, pubkeys [][]byte) (string, error) {
	if n > 16 || n < 1 {
		return "", fmt.Errorf("invalid n value")
	}
	if n > m || m < 1 {
		return "", fmt.Errorf("invalid m value")
	}
	if len(pubkeys) != m {
		return "", fmt.Errorf("invalid length of public keys")
	}

	builder := txscript.NewScriptBuilder()
	// add the minimum number of needed signatures
	builder.AddOp(byte(0x50 + n))
	// add the public keys
	for _, pk := range pubkeys {
		builder.AddData(pk)
	}
	// add the total number of public keys in the multi-sig
	builder.AddOp(byte(0x50 + m))
	// add the check-multi-sig op-code
	builder.AddOp(txscript.OP_CHECKMULTISIG)

	redeemScript, err := builder.Script()
	if err != nil {
		return "", err
	}

	redeemHash := btcutil.Hash160(redeemScript)
	address, err := btcutil.NewAddressScriptHashFromHash(redeemHash, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	return address.EncodeAddress(), nil
}
