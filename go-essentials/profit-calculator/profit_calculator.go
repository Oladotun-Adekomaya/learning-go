package main

import (
	"errors"
	"fmt"
	// "math"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	// fmt.Println(revenue, expenses, taxRate)

	// ebt := revenue - expenses
	// profit := ebt - (ebt*(taxRate/100))
	// ratio := ebt / profit

	// fmt.Printf("The earnings before tax and interest is: %v\n", ebt)
	// fmt.Printf("The profit is %v\n",profit)
	// fmt.Printf("The ratio is %.2f\n",ratio)

	// ratiotxt := fmt.Sprintf("The earnings to profit ratio is %.2f\n",ratio)
	// profittxt := fmt.Sprintf("The profit is %v\n",profit)
	// earningstxt := fmt.Sprintf("The earnings before tax and interest is: %v\n", ebt)

	// fmt.Print(earningstxt, profittxt, ratiotxt)

	//FUNCTIONS
	revenue, err := getInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculations(revenue, expenses, taxRate)

	fmt.Printf("%.1f, %.1f, %.2f\n", ebt, profit, ratio)

}

func getInput(text string) (float64, error) {
	fmt.Printf("%v", text)
	var value float64
	fmt.Scanln(&value)

	if value <= 0 {
		return 0, errors.New("Value must be positive")
	}
	return value, nil
}

func calculations(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (ebt * (taxRate / 100))
	ratio = ebt / profit
	return
}
