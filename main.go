package main

import (
	"fmt"
	"github.com/Mokyton/DeNET/account"
	"log"
	"os"
)

func main() {
	var password []byte
	var login []byte
	acc := account.New()
	if isAccountExist() {
		greetingToAccess(password, login)
		res, err := acc.SignIn(password, login)
		if err != nil {
			log.Fatal(err)
		}
		if !res {
			fmt.Println("--- Wrong login or password")
			fmt.Println("--- Please try again later")
			os.Exit(0)
		}
		fmt.Println("--- Congratulations you are logged in!")
		fmt.Println("--- This is your public key:")
		fmt.Println(acc.PubKey)
	} else {
		greetingCreateAcc(password)
		err := acc.CreateAccount(password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("--- This is your wallet address! Use it as login next time:")
		fmt.Println(acc.Address)
		fmt.Println("--- This is your public key:")
		fmt.Println(acc.PubKey)
	}

}
