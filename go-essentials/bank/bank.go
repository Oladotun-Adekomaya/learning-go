package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {

}

func main() {
	var accountBalance float64 = 1000
	fmt.Println("Welcome to the Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			fmt.Printf("Balance updated! New amount: $%v\n", accountBalance)

		} else {
			println("Goodbye")
			break
		}
	}

	fmt.Println("This the end the code")
	// fmt.Println("Your choice:",choice)
}
