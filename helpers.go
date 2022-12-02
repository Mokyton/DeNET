package main

import (
	"bytes"
	"fmt"
	"os"
)

func greetingCreateAcc(password []byte) {
	var choice int
	var confirmPs []byte
	counter := 0
	fmt.Println("--- Hello, Sorry you don't have an account")
	fmt.Println("--- Press: \n--- 1 to create account\n--- 2 to exit")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("--- Enter the password")
		fmt.Println("--- Password shouldn't be less than 8 symbols")
		fmt.Scan(&password)
		for len(string(password)) < 8 {
			fmt.Println("--- Password less than 8 symbols, please try new one")
			fmt.Scan(&password)
			counter++
			if counter == 3 {
				fmt.Println("--- Sorry password isn't valid. Come and try later")
				os.Exit(0)
			}
		}

		fmt.Println("--- Please confirm your password ---")
		fmt.Scan(&confirmPs)
		if bytes.Compare(password, confirmPs) != 0 {
			fmt.Println("--- Error: different passwords")
			fmt.Println("--- Try again later ---")
			os.Exit(0)
		}
	} else {
		os.Exit(0)
	}

}

func greetingToAccess(password []byte, login []byte) {
	fmt.Println("--- Hello, you already have an account")

	fmt.Println("--- Please enter your login")
	fmt.Scan(&login)

	fmt.Println("--- Please enter your password")
	fmt.Scan(&password)
}

func isAccountExist() bool {
	if _, err := os.Stat("./storage/wallets"); os.IsNotExist(err) {
		return false
	}
	return true
}
