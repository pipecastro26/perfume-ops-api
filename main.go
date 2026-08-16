package main

import "fmt"

func main() {
	const currency = "COP"
	const minimumStock = 3

	productName := "Yara"
	brand := "Lattafa"
	purchasePrice := 100000
	salePrice := 200000
	stock := 8
	soldUnits := 2
	available := true

	remainingStock := stock - soldUnits
	profitPerUnit := salePrice - purchasePrice
	totalRevenue := salePrice * soldUnits
	totalProfit := profitPerUnit * soldUnits
	lowStock := remainingStock <= minimumStock

	fmt.Println("PerfumeOps")
	fmt.Println("--------------------")
	fmt.Printf("Product: %s\n", productName)
	fmt.Printf("Brand: %s\n", brand)
	fmt.Printf("Sale price: %d %s\n", salePrice, currency)
	fmt.Printf("Units sold: %d\n", soldUnits)
	fmt.Printf("Remaining stock: %d\n", remainingStock)
	fmt.Printf("Revenue: %d %s\n", totalRevenue, currency)
	fmt.Printf("Profit per unit: %d %s\n", profitPerUnit, currency)
	fmt.Printf("Total profit: %d %s\n", totalProfit, currency)
	fmt.Printf("Low stock: %t\n", lowStock)
	if remainingStock == 0 {
		fmt.Println("Inventory status: Out of stock")
	} else if remainingStock <= minimumStock {
		fmt.Println("Inventory status: Low stock")
	} else {
		fmt.Println("Inventory status: Sufficient")
	}
	if available && remainingStock > 0 {
		fmt.Println("Product can be sold")
	} else {
		fmt.Println("Product cannot be sold")
	}

}
