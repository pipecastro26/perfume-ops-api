package main

import (
	"errors"
	"fmt"
)

func calculateDiscount(subtotal, porcentaje_de_descuento int) int {
	return (subtotal * porcentaje_de_descuento) / 100
}
func calculateFinalTotal(subtotal int, porcentaje_de_descuento int) int {
	return subtotal - ((subtotal * porcentaje_de_descuento) / 100)
}
func processSale(stock, quantity int) (int, error) {
	if quantity <= 0 {
		return 0, errors.New("quantity must be greater than zero")
	} else if quantity > stock {
		return 0, errors.New("insufficient stock")
	} else {
		return stock - quantity, nil
	}
}
func main() {
	subtotal := 600000
	discount := 10
	stock := 10
	quantity := 0
	fmt.Println("Descuento: ", calculateDiscount(subtotal, discount))
	fmt.Println("Total", calculateFinalTotal(subtotal, discount))
	remainingStock, err := processSale(stock, quantity)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Stock: ", remainingStock)

}
