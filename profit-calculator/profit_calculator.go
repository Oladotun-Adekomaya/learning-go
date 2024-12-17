package main

import (
	"fmt"
	// "math"
	
)

func main(){
	var revenue float64
	var expenses float64
	var taxRate float64

	
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// fmt.Println(revenue, expenses, taxRate)

	ebt := revenue - expenses
	profit := ebt - (ebt*(taxRate/100))
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	
}