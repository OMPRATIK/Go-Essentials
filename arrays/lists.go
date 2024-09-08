package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [3]string
	prices := [4]float64{1, 2.8, 3, 4.1}

	productNames[2] = "Pratik"

	// featuredPrices := prices[1 : 4]
	// featuredPrices := prices[: 3] // starts with 0
	featuredPrices := prices[1 :] // starts with 1 ends with 4 (exclusive)
	
	fmt.Println(prices)
	fmt.Println(prices[1])
	fmt.Println(productNames)

	fmt.Println("Slice of an array", featuredPrices)
}