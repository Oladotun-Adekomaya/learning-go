package main

import(
	"fmt"
)

func main(){
	var accountBalance float64 = 1000
	

	fmt.Println("Welcome to the Go Bank!")
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
	}else if choice == 2 {

		fmt.Println("Your deposit: ")
		var deposit float64
		fmt.Scan(&deposit)
		accountBalance += deposit

	}else if choice == 3 {

		fmt.Println("Your Withdrawal: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount

	}

	// fmt.Println("Your choice:",choice)
}