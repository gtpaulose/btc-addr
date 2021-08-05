package keystore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const mnemonic = "aim option gasp apple cabbage film pig soldier wine pact spot book"

func TestPrintAddressForPathPurposeBIP44(t *testing.T) {
	ks := New(WithMnemonic(mnemonic))
	address := ks.printAddressForPath(PurposeBIP44, GetPath(PurposeBIP44, 0))

	assert.Equal(t, address.Get(), "1Dkga7whWV3vH1pRxNwTfa6TFsFasgQeLF")
	assert.Equal(t, fmt.Sprintf("%x", address.GetKey().GetPublicKey()), "03810329392848b317c942580ee9c71e2d0f478d99670a420d61c4fd63a1919c26")
	assert.Equal(t, address.GetKey().GetWIF(), "L5SjSCBwJQJmoBpfTVxDn1v5vKRJQXSmByoHAKyXTAP6wKu4Ncxs")
}

func TestPrintAddressForPathPurposeBIP49(t *testing.T) {
	ks := New(WithMnemonic(mnemonic))
	address := ks.printAddressForPath(PurposeBIP49, GetPath(PurposeBIP49, 0))

	assert.Equal(t, address.Get(), "36JrL92NvxbeuSFRggfWVkmtk1bgfwgYey")
	assert.Equal(t, fmt.Sprintf("%x", address.GetKey().GetPublicKey()), "0395096ac06b4d80943a1a93b09eb82caa6b72f3b76fab10a077de5c635f05c03d")
	assert.Equal(t, address.GetKey().GetWIF(), "L4yeb4WMRbaJMqAFKEDhU3AY81Q4RNtfDMdDvze9W4iJwi3d7guj")
}

func TestPrintAddressForPathPurposeBIP84(t *testing.T) {
	ks := New(WithMnemonic(mnemonic))
	address := ks.printAddressForPath(PurposeBIP84, GetPath(PurposeBIP84, 0))

	assert.Equal(t, address.Get(), "bc1qtszqcgzu84alenagf4h3g3x5ewxylgy6ze08m9")
	assert.Equal(t, fmt.Sprintf("%x", address.GetKey().GetPublicKey()), "0398781494367e40b9d4cfe99d420986a022d23a89f3bfdca62c4c5f4ab46556b9")
	assert.Equal(t, address.GetKey().GetWIF(), "KxUZEVFxo9KfgQmm9BZbEwPZR1y3kJPbNXWtbFj1CXdb4T1dmu9N")
}
