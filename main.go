package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/Mokyton/DeNET/hardcodeHash"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"log"
	"os"
)

type account struct {
	address  string
	password string
}

type wallet struct {
}

func main() {
	acc, err := CreateAccount()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(acc)
}

func CreateAccount() (account, error) {
	var password []byte
	fmt.Scan(&password)
	pass := sha256.Sum256(password)
	passHash := fmt.Sprintf("%x\n", pass)
	encodedPassHash, err := hardcodeHash.Encrypt(hardcodeHash.KEY, []byte(passHash))
	file, err := os.Create("hardcodedPass/encodedPassHash.txt")
	if err != nil {
		return account{}, err
	}
	_, err = file.Write(encodedPassHash)
	if err != nil {
		return account{}, err
	}
	address, err := CreateWallet(passHash)
	if err != nil {
		return account{}, err
	}
	return account{address: address, password: passHash}, nil
}

func CreateWallet(passphrase string) (string, error) {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(passphrase)
	if err != nil {
		return "", err
	}
	return account.Address.Hex(), nil
}
