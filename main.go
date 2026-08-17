package main

import (
	"errors"
	"fmt"
)

func calculateProfit(salePrice, purchasePrice int) int {
	return salePrice - purchasePrice
}

func calculateRevenue(price, quantity int) int {
	return price * quantity
}

func isLowStock(stock, minimumStock int) bool {
	return stock <= minimumStock
}

func isAvailable(stock int) bool {
	return stock > 0
}

func processSale(stock, quantity int) (int, error) {
	if quantity <= 0 {
		return 0, errors.New("quantity must be greater than zero")
	}

	if quantity > stock {
		return 0, errors.New("insufficient stock")
	}

	return stock - quantity, nil
}

func main() {
	const currency = "COP"
	const minimumStock = 3

	productName := "Yara"
	brand := "Lattafa"
	purchasePrice := 100000
	salePrice := 200000
	stock := 8
	soldUnits := 2

	remainingStock, err := processSale(stock, soldUnits)

	if err != nil {
		fmt.Println("Sale could not be processed:", err)
		return
	}

	profitPerUnit := calculateProfit(salePrice, purchasePrice)
	totalRevenue := calculateRevenue(salePrice, soldUnits)
	totalProfit := profitPerUnit * soldUnits
	lowStock := isLowStock(remainingStock, minimumStock)
	available := isAvailable(remainingStock)

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
	fmt.Printf("Available: %t\n", available)

	if remainingStock == 0 {
		fmt.Println("Inventory status: Out of stock")
	} else if lowStock {
		fmt.Println("Inventory status: Low stock")
	} else {
		fmt.Println("Inventory status: Sufficient")
	}
}
