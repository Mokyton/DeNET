package account

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/Mokyton/DeNET/cipherHash"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"os"
)

type account struct {
	Address string
	PubKey  string
}

func New() account {
	return account{}
}

func (acc *account) CreateAccount(password []byte) error {
	pass := sha256.Sum256(password)
	passPhrase := fmt.Sprintf("%x", pass)

	log := sha256.Sum256([]byte(acc.Address))
	logPhrase := fmt.Sprintf("%x", log)

	err := acc.savePassHash([]byte(passPhrase))
	if err != nil {
		return err
	}

	err = acc.CreateWallet(passPhrase)
	if err != nil {
		return err
	}

	err = acc.saveLoginHash([]byte(logPhrase))
	if err != nil {
		return err
	}

	return nil
}

func (acc *account) CreateWallet(passphrase string) error {
	ks := keystore.NewKeyStore("./storage/wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.NewAccount(passphrase)
	if err != nil {
		return err
	}
	err = acc.getPublicKeyAndAddress(passphrase)
	if err != nil {
		return err
	}

	return nil
}

func (acc *account) SignIn(password []byte, login []byte) (bool, error) {
	cL, err := acc.checkLogin(login)
	if err != nil {
		return false, err
	}
	cP, err := acc.checkPassword(password)
	if err != nil {
		return false, err
	}

	pass := sha256.Sum256(password)
	passPhrase := fmt.Sprintf("%x", pass)

	if !cL || !cP {
		return false, nil
	}

	if err = acc.getPublicKeyAndAddress(passPhrase); err != nil {
		return false, err
	}

	return true, nil
}

func (acc *account) getPublicKeyAndAddress(passphrase string) error {
	files, err := ioutil.ReadDir("./storage/wallets")
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile("storage/wallets/" + files[0].Name())
	if err != nil {
		return err
	}

	key, err := keystore.DecryptKey(b, passphrase)
	if err != nil {
		return err
	}

	pk := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	acc.PubKey = hexutil.Encode(pk)[2:]
	acc.Address = crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()[2:]

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

func (acc *account) checkPassword(enteredPass []byte) (bool, error) {
	data, err := ioutil.ReadFile("./storage/accountHash/encodedPassHash.txt")
	if err != nil {
		return false, err
	}

	passFromDB, err := cipherHash.Decrypt(cipherHash.KEY, data)
	if err != nil {
		return false, err
	}

	enteredHash := sha256.Sum256(enteredPass)
	passHex := fmt.Sprintf("%x", enteredHash)

	return bytes.Compare(passFromDB, []byte(passHex)) == 0, nil

}

func (acc *account) checkLogin(enteredLogin []byte) (bool, error) {
	data, err := ioutil.ReadFile("./storage/accountHash/encodedLoginHash.txt")
	if err != nil {
		return false, err
	}

	loginFromDB, err := cipherHash.Decrypt(cipherHash.KEY, data)
	if err != nil {
		return false, err
	}

	enteredHash := sha256.Sum256(enteredLogin)
	loginHex := fmt.Sprintf("%x", enteredHash)

	return bytes.Compare(loginFromDB, []byte(loginHex)) == 0, nil
}
