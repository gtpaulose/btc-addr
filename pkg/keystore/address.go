package keystore

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

type Address interface {
	Generate() error
	Get() string
	GetKey() Keyer
}

type BIP44 struct {
	k       *Key
	address string
}

func (a *BIP44) Get() string   { return a.address }
func (a *BIP44) GetKey() Keyer { return a.k }

func (a *BIP44) Generate() error {
	addressPubKey, err := btcutil.NewAddressPubKey(a.k.wif.SerializePubKey(), &chaincfg.MainNetParams)
	if err != nil {
		return err
	}
	a.address = addressPubKey.EncodeAddress()

	return nil
}

type BIP49 struct {
	k            *Key
	segwitNested string
}

func (a *BIP49) Get() string   { return a.segwitNested }
func (a *BIP49) GetKey() Keyer { return a.k }

func (a *BIP49) Generate() error {
	witnesshash, err := getWitnessHash(a.k.wif.SerializePubKey())
	if err != nil {
		return err
	}

	script, err := txscript.PayToAddrScript(witnesshash)
	if err != nil {
		return err
	}
	scriptHash, err := btcutil.NewAddressScriptHash(script, &chaincfg.MainNetParams)
	if err != nil {
		return err
	}
	a.segwitNested = scriptHash.EncodeAddress()

	return nil
}

type BIP84 struct {
	k            *Key
	segwitBech32 string
}

func (a *BIP84) Get() string   { return a.segwitBech32 }
func (a *BIP84) GetKey() Keyer { return a.k }

func (a *BIP84) Generate() error {
	witnesshash, err := getWitnessHash(a.k.wif.SerializePubKey())
	if err != nil {
		return err
	}

	a.segwitBech32 = witnesshash.EncodeAddress()

	return nil
}

func getWitnessHash(pubkey []byte) (*btcutil.AddressWitnessPubKeyHash, error) {
	witness := btcutil.Hash160(pubkey)
	return btcutil.NewAddressWitnessPubKeyHash(witness, &chaincfg.MainNetParams)
}
