package account

import (
	"crypto/sha256"
	"fmt"
	"github.com/Mokyton/DeNET/account/wallet"
	"github.com/Mokyton/DeNET/cipherHash"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"os"
)

type account struct {
	Login    string
	Password string
	Wallet   wallet.Ks
}

func New() account {
	ks := wallet.New()
	return account{Wallet: ks}
}

func (acc *account) CreateAccount(password []byte) error {
	pass := sha256.Sum256(password)
	acc.Password = fmt.Sprintf("%x", pass)

	err := acc.savePassHash([]byte(acc.Password))
	if err != nil {
		return err
	}

	acc.Login, err = acc.Wallet.CreateWallet(acc.Password)
	if err != nil {
		return err
	}

	err = acc.saveLoginHash([]byte(acc.Login))
	if err != nil {
		return err
	}

	return nil
}

func (acc *account) GetPublicKeyAndAddress() error {
	files, err := ioutil.ReadDir("./storage/wallets")
	if err != nil {
		return err
	}
	fmt.Println(files[0].Name())
	b, err := ioutil.ReadFile(files[0].Name())
	if err != nil {
		return err
	}

	key, err := keystore.DecryptKey(b, acc.Password)
	if err != nil {
		return err
	}

	pk := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	acc.Wallet.PublicKey = hexutil.Encode(pk)
	acc.Wallet.Address = crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println(acc.Wallet.Address)
	fmt.Println(acc.Wallet.PublicKey)
	return nil
}

func (acc *account) savePassHash(passHash []byte) error {
	encodedPassHash, err := cipherHash.Encrypt(cipherHash.KEY, passHash)
	if err != nil {
		return err
	}

	file, err := os.Create("storage/accountHash/encodedPassHash.txt")
	_, err = file.Write(encodedPassHash)
	if err != nil {
		return err
	}

	return nil
}

func (acc *account) saveLoginHash(loginHash []byte) error {
	encodedLoginHash, err := cipherHash.Encrypt(cipherHash.KEY, loginHash)
	if err != nil {
		return err
	}

	file, err := os.Create("storage/accountHash/encodedLoginHash.txt")
	_, err = file.Write(encodedLoginHash)
	if err != nil {
		return err
	}

	return nil
}
