package keystore

import (
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/tyler-smith/go-bip32"
)

const (
	compress = true
)

type Keyer interface {
	GetWIF() string
	GetPublicKey() []byte
	GenerateWIF() (*btcutil.WIF, error)
}

type Key struct {
	*bip32.Key
	pk   *btcec.PrivateKey
	pubk *btcec.PublicKey
	wif  *btcutil.WIF
}

func NewKey(k *bip32.Key) *Key {
	pk, pubk := btcec.PrivKeyFromBytes(btcec.S256(), k.Key)
	key := &Key{k, pk, pubk, nil}
	wif, err := key.GenerateWIF()
	if err != nil {
		log.Fatalln("Error generating WIF: ", err)
	}
	key.wif = wif

	return key
}

func (k *Key) GenerateWIF() (*btcutil.WIF, error) {
	return btcutil.NewWIF(k.pk, &chaincfg.MainNetParams, compress)
}

func (k *Key) GetWIF() string       { return k.wif.String() }
func (k *Key) GetPublicKey() []byte { return k.pubk.SerializeCompressed() }
