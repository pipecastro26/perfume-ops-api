package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	SKU   string
	Name  string
	Price int
	Stock int
}

func main() {
	products := []Product{
		{
			SKU:   "A",
			Name:  "Asad",
			Price: 200000,
			Stock: 3,
		},
		{
			SKU:   "B",
			Name:  "BORB",
			Price: 300000,
			Stock: 4,
		},
		{
			SKU:   "C",
			Name:  "Yara",
			Price: 400000,
			Stock: 9,
		},
	}
	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	err = os.WriteFile("practiceProducts.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Products saved successfully")

	dataFromFile, err := os.ReadFile("practiceProducts.json")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var loadedProducts []Product

	err = json.Unmarshal(dataFromFile, &loadedProducts)

	if err != nil {
		fmt.Println("Error converting JSON:", err)
		return
	}

	fmt.Println()
	fmt.Println("Products loaded:")

	for _, product := range loadedProducts {
		fmt.Println("--------------------")
		fmt.Println("SKU:", product.SKU)
		fmt.Println("Name:", product.Name)
		fmt.Println("Price:", product.Price)
		fmt.Println("Stock:", product.Stock)
	}

}
