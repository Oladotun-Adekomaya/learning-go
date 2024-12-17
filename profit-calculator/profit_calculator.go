package main

import (
	"fmt"
	// "math"
	
)

func main(){
	var revenue, expenses, taxRate float64
	
	fmt.Println("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Println("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Println("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt - (ebt*(taxRate/100))
	ratio := ebt/profit

	println("Ebt: ", ebt)
	println("Expenses: ", expenses)
	println("Ratio: ", ratio)
}