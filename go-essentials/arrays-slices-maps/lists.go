package main

import (
	"fmt"
)

func main() {
	prices := []float64{10.99, 8.99}

	prices = append(prices, 5.99)

	fmt.Println(prices)
}

// func main() {
// 	prices := [4]float64{10.99, 30.0, 19.5, 50.0}

// 	var productNames [3]string = [3]string{"A Book", "A Pen"}

// 	productNames[2] = "A carpet"
// 	fmt.Println(productNames)
// 	fmt.Println(prices[0])

// 	featuredPrices := prices[1:3]

// 	featuredPrices[0] = 199.99

// 	// fmt.Println(featuredPrices)
// 	fmt.Println(len(prices), cap(prices))
// }
