package main

import (
	"fmt"
	"os"
	"strconv"
)

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile("balance.txt")
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}

func writeBalanceToFile(balance float64) {
	balancetxt := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balancetxt), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile()
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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
			fmt.Printf("Balance updated! New amount: $%v\n", accountBalance)

		} else {
			println("Goodbye")
			break
		}
	}

	fmt.Println("This the end the code")
	// fmt.Println("Your choice:",choice)
}
