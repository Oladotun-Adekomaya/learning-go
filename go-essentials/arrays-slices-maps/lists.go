package main

import (
	"fmt"
)

func main() {
	prices := [4]float64{10.99, 30.0, 19.5, 50.0}

	var productNames [3]string = [3]string{"A Book", "A Pen"}

	productNames[2] = "A carpet"
	fmt.Println(productNames)
	fmt.Println(prices[0])

	featuredPrices := prices[1:3]

	fmt.Println(featuredPrices)
}
