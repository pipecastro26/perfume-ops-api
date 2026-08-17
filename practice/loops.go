package main

import "fmt"

const moneda string = "COP"

func main() {
	priceUnit := 210000
	numberofSale := 4
	total := 0

	for i := 0; i < numberofSale; i++ {
		fmt.Println("Sale ", i, ": ", priceUnit, moneda)
		total = priceUnit + total
	}
	fmt.Print("Total: ", total, moneda)

}
