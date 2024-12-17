package main

import (
	"fmt"
	"math"
)

func main()  {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64   

	fmt.Print("Enter Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Years: ")
	fmt.Scan(&years)

	fmt.Print("Enter Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100,years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue, futureRealValue)
}