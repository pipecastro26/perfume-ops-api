package main

import (
	"errors"
	"fmt"
)

func calculateRevenue(price, quantity int) int {
	return price * quantity
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

func displayProduct(index int, name string, price int, stock int, minimumStock int) {
	fmt.Printf(
		"%d. %s - %d COP - Stock: %d",
		index+1,
		name,
		price,
		stock,
	)

	if stock <= minimumStock {
		fmt.Print(" - WARNING: Low stock")
	}

	fmt.Println()
}

func main() {
	const currency = "COP"
	const minimumStock = 5

	products := []string{
		"Yara",
		"Khamrah",
		"Asad Bourbon",
	}

	prices := []int{
		200000,
		210000,
		210000,
	}

	stocks := []int{
		8,
		4,
		10,
	}

	fmt.Println("PerfumeOps Catalog")
	fmt.Println("--------------------")

	totalInventoryValue := 0

	for index, product := range products {
		displayProduct(
			index,
			product,
			prices[index],
			stocks[index],
			minimumStock,
		)

		totalInventoryValue += prices[index] * stocks[index]
	}

	fmt.Println("--------------------")
	fmt.Printf("Total inventory value: %d %s\n", totalInventoryValue, currency)

	// Primera venta
	productIndex := 0
	quantity := 3

	fmt.Println()
	fmt.Println("Processing sale...")
	fmt.Println("--------------------")

	remainingStock, err := processSale(
		stocks[productIndex],
		quantity,
	)

	if err != nil {
		fmt.Println("Sale could not be processed:", err)
	} else {
		stocks[productIndex] = remainingStock

		saleRevenue := calculateRevenue(
			prices[productIndex],
			quantity,
		)

		fmt.Println("Sale processed successfully")
		fmt.Println("Product:", products[productIndex])
		fmt.Println("Units sold:", quantity)
		fmt.Println("Remaining stock:", stocks[productIndex])
		fmt.Printf("Sale revenue: %d %s\n", saleRevenue, currency)
	}

	// Segunda venta
	productIndex2 := 1
	quantity2 := 3

	fmt.Println()
	fmt.Println("Processing second sale...")
	fmt.Println("--------------------")

	remainingStock2, err := processSale(
		stocks[productIndex2],
		quantity2,
	)

	if err != nil {
		fmt.Println("Sale could not be processed:", err)
	} else {
		stocks[productIndex2] = remainingStock2

		saleRevenue2 := calculateRevenue(

			prices[productIndex2],
			quantity2,
		)

		fmt.Println("Sale processed successfully")
		fmt.Println("Product:", products[productIndex2])
		fmt.Println("Units sold:", quantity2)
		fmt.Println("Remaining stock:", stocks[productIndex2])
		fmt.Printf("Sale revenue: %d %s\n", saleRevenue2, currency)
	}
}
