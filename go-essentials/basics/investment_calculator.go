package main

import (
	"fmt"
	"math"
)
const inflationRate = 2.5
func main()  {
	
	var investmentAmount, years, expectedReturnRate float64   

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100,years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue, futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64){
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100,years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}