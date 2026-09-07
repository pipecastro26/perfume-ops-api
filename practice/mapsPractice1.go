package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func (p *Product) IncreaseStock(quantity int) error {
	if quantity <= 0 {
		return errors.New("Error")
	}
	p.Stock += quantity
	return nil
}
func saveProducts(products []Product, filename string) error {
	data, err := json.MarshalIndent(products, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		return err
	}

	return nil
}
func loadProducts(filename string) ([]Product, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return []Product{}, errors.New("no se leyo")
	}

	var products []Product

	err = json.Unmarshal(data, &products)

	if err != nil {
		return products, errors.New("no se leyo")
	}

	return products, nil
}
func main() {
	products := []Product{
		{Name: "Yara", Price: 200000, Stock: 8},
		{Name: "Khamrah", Price: 210000, Stock: 4},
		{Name: "Asad Bourbon", Price: 210000, Stock: 10},
	}
	err := saveProducts(products, "products.json")
	prop, err2 := loadProducts("products.json")
	if err2 != nil {
		fmt.Println("Error saving products:", err2)
	} else {
		fmt.Println(prop)
	}
	if err != nil {
		fmt.Println("Error saving products:", err)
	} else {
		fmt.Println("Products saved successfully")
	}
}
