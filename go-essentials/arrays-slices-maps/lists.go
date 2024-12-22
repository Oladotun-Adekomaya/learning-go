package main

import (
	"fmt"
)

func main() {
	prices := [4]float64{10.99, 30.0, 19.5, 50.0}

	var productNames [3]string = [3]string{"A Book", "A Pen"}

	fmt.Println(productNames)
	fmt.Println(prices[0])
}
