package wallet

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type Ks struct {
	PublicKey string
	Address   string
	cd        bool
}

func New() Ks {
	return Ks{}
}

func (w *Ks) CreateWallet(passphrase string) (string, error) {
	ks := keystore.NewKeyStore("./storage/wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(passphrase)
	if err != nil {
		return "", err
	}
	return account.Address.Hex(), nil
}
