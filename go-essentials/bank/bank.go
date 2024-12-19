package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		// panic("Cant continue, sorry")
	}

	fmt.Println("Welcome to the Go Bank!")

	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your account balance is $%v\n", accountBalance)
		} else if choice == 2 {

			fmt.Print("Your deposit: ")
			var deposit float64
			fmt.Scan(&deposit)

			if deposit < 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += deposit
			writeFloatToFile(accountBalanceFile, accountBalance)
			fmt.Printf("Balance updated! New amount: $%v\n", accountBalance)

		} else if choice == 3 {

			fmt.Print("Withdrawal Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount < 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Must not be more than account balance.")
				continue
			}
			accountBalance -= withdrawAmount
			writeFloatToFile(accountBalanceFile, accountBalance)
			fmt.Printf("Balance updated! New amount: $%v\n", accountBalance)

		} else {
			println("Goodbye")
			break
		}
	}

	fmt.Println("This the end the code")
	// fmt.Println("Your choice:",choice)
}
