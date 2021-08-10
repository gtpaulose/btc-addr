package keystore

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cast"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type Purpose int

const (
	apostrophe uint32 = 0x80000000 // 0'

	PurposeBIP44 Purpose = 44 // BIP44
	PurposeBIP49 Purpose = 49 // BIP49
	PurposeBIP84 Purpose = 84 // BIP84
)

type KeyStore struct {
	mnemonic   string
	passphrase string
}

type KeyStoreOption func(*KeyStore)

func WithPassphrase(passphrase string) KeyStoreOption {
	return func(ks *KeyStore) {
		ks.passphrase = passphrase
	}
}

func WithMnemonic(mnemonic string) KeyStoreOption {
	return func(ks *KeyStore) {
		ks.mnemonic = mnemonic
	}
}

func GenerateMnemonic(phrases int) KeyStoreOption {
	return func(ks *KeyStore) {
		mnemonic, err := generateMnemonic(phrases)
		if err != nil {
			log.Fatalln("Error generating mnemonic: ", err)
		}

		ks.mnemonic = mnemonic
	}
}

func New(opts ...KeyStoreOption) *KeyStore {
	ks := &KeyStore{}
	for _, o := range opts {
		o(ks)
	}

	return ks
}

func generateMnemonic(phrases int) (string, error) {
	entropy, err := bip39.NewEntropy(phrases * 32 / 3)
	if err != nil {
		return "", err
	}

	return bip39.NewMnemonic(entropy)
}

func (ks *KeyStore) GetMnemonic() string { return ks.mnemonic }

func (ks *KeyStore) GetSeed() string {
	return string(bip39.NewSeed(ks.mnemonic, ks.passphrase))
}

func (ks *KeyStore) GetRootKey() *bip32.Key {
	key, err := bip32.NewMasterKey([]byte(ks.GetSeed()))
	if err != nil {
		fmt.Println("Error getting master key")
		return nil
	}

	return key
}

func (ks *KeyStore) Print() {
	fmt.Printf("%-10s %s\n", "Mnemonic:", ks.mnemonic)
	fmt.Printf("%-10s %s\n", "PassPhrase:", ks.passphrase)
	fmt.Printf("%-10s %x\n", "Seed:", ks.GetSeed())
	fmt.Printf("%-10s %s\n", "Root Key (BIP32):", ks.GetRootKey().B58Serialize())
}

func GetPath(purpose Purpose, index int) string {
	return fmt.Sprintf(`m/%d'/0'/0'/0/%d`, purpose, index)
}

func getKey(children []string, parent *bip32.Key) (*bip32.Key, error) {
	child, err := parent.NewChildKey(getChildIndex(children[0]))
	if err != nil {
		return nil, err
	}

	if len(children) == 1 {
		return child, nil
	}

	return getKey(children[1:], child)
}

func getChildIndex(index string) uint32 {
	if index[len(index)-1] == 39 {
		return cast.ToUint32(index[:len(index)-1]) + apostrophe
	}

	return cast.ToUint32(index)
}

func (ks *KeyStore) PrintAddresses(purpose Purpose, n int) {
	printTable()
	for i := 0; i < n; i++ {
		path := GetPath(purpose, i)
		ks.printAddressForPath(purpose, path)
	}
}

func (ks *KeyStore) PrintAddressForPath(purpose Purpose, path string) {
	printTable()
	ks.printAddressForPath(purpose, path)
}

func (ks *KeyStore) printAddressForPath(purpose Purpose, path string) Address {
	root := ks.GetRootKey()
	key, err := getKey(strings.Split(path, "/")[1:], root)
	if err != nil {
		fmt.Println("error getting key: ", err)
		return nil
	}

	var a Address
	switch purpose {
	case PurposeBIP44:
		a = &BIP44{
			k: NewKey(key),
		}
	case PurposeBIP49:
		a = &BIP49{
			k: NewKey(key),
		}
	case PurposeBIP84:
		a = &BIP84{
			k: NewKey(key),
		}
	default:
		fmt.Println("invalid purpose: ", err)
		return nil
	}

	if err := a.Generate(); err != nil {
		fmt.Println("error generating address: ", err)
		return nil
	}

	fmt.Printf("%-18s %-34s %-52x %50s\n", path, a.Get(), a.GetKey().GetPublicKey(), a.GetKey().GetWIF())
	return a
}

func printTable() {
	fmt.Println(strings.Repeat("-", 173))
	fmt.Printf("\n%-18s %-34s %-52s %25s\n", "Path", "Address", "Public Key", "Private Key")
	fmt.Println(strings.Repeat("-", 173))
}

func ValidateFormat(format string) (Purpose, error) {
	switch format {
	case "bip44":
		return PurposeBIP44, nil
	case "bip49":
		return PurposeBIP49, nil
	case "bip84":
		return PurposeBIP84, nil
	default:
		return Purpose(-1), fmt.Errorf("invalid format")
	}
}

func ValidatePath(path string, purpose Purpose) error {
	if len(strings.Split(path, "/")) > 2 {
		index := strings.Split(path, "/")[1]
		if getChildIndex(index) != uint32(purpose)+apostrophe {
			return fmt.Errorf("inconsistent path and format")
		}
	}

	return nil
}
