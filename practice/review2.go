package main

import (
	"errors"
	"fmt"
)

func processOrder(stock, quantity int) (int, error) {
	if quantity <= 0 {
		return 0, errors.New("quantity must be greater than zero")
	}

	if quantity > stock {
		return 0, errors.New("insufficient stock")
	}

	return stock - quantity, nil
}

func main() {
	stock := 3
	quantity := 8

	remainingStock, err := processOrder(stock, quantity)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Remaining stock:", remainingStock)
}
