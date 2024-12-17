package main

import (
	"math"
	"fmt"
)

func main()  {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate,float64(years))

	fmt.Print(futureValue)
}