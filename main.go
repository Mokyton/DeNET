package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/Mokyton/DeNET/cipherHash"
	"io/ioutil"
	"os"
)

func main() {
	//acc := account.New()
	//if isAccountExist() {
	//
	//} else {
	//	err := acc.CreateAccount([]byte("CRINGE"))
	//	_ = acc.GetPublicKeyAndAddress()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	fmt.Println(checkPassword([]byte("CRINGE")))
}

func checkPassword(enteredPass []byte) (bool, error) {
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

//files, err := ioutil.ReadDir("./storage/wallets")
//if err != nil {
//log.Fatal(err)
//}
//
//fmt.Println(files[0].Name())

func isAccountExist() bool {
	if _, err := os.Stat("./storage/wallets"); os.IsNotExist(err) {
		return false
	}
	return true
}
