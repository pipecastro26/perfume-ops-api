package main

import "fmt"

func main() {
	productName := "Khamrah"
	brand := "Lattafa"

	purchasePrice := 110000
	salePrice := 210000

	stock := 5
	soldUnits := 2

	const minimumStock = 3
	const currency = "COP"

	profitPerUnit := salePrice - purchasePrice

	remainingStock := stock - soldUnits

	totalRevenue := salePrice * soldUnits

	totalProfit := profitPerUnit * soldUnits

	lowStock := remainingStock <= minimumStock

	fmt.Printf("Product: %s\n", productName)
	fmt.Printf("Brand: %s\n", brand)
	fmt.Printf("Remaining stock: %d\n", remainingStock)
	fmt.Printf("Total revenue: %d %s\n", totalRevenue, currency)
	fmt.Printf("Total profit: %d %s\n", totalProfit, currency)
	fmt.Printf("Low stock: %t\n", lowStock)
}
