package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/pipecastro26/perfume-ops-api/internal/domain"
	"github.com/pipecastro26/perfume-ops-api/internal/inventory"
)

func displayProduct(index int, product domain.Product, minimumStock int) {
	fmt.Printf("%d. %s - %s - %d COP - Stock: %d", index+1, product.SKU, product.Name, product.Price, product.Stock)

	if product.Stock <= minimumStock {
		fmt.Print(" - WARNING: Low stock")
	}

	fmt.Println()
}

func saveProducts(products []domain.Product, filename string) error {
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

func loadProducts(filename string) ([]domain.Product, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return []domain.Product{}, errors.New("could not read products file")
	}

	var products []domain.Product

	err = json.Unmarshal(data, &products)

	if err != nil {
		return products, errors.New("could not decode products JSON")
	}

	return products, nil
}

func printSale(products []domain.Product, quantity int, index int, err error, revenue int) {
	if err != nil {
		fmt.Println("Sale could not be processed:", err)
		return
	}

	fmt.Println("Sale processed successfully")
	fmt.Println("Product:", products[index].Name)
	fmt.Println("Units sold:", quantity)
	fmt.Println("Remaining stock:", products[index].Stock)
	fmt.Println("Revenue:", revenue)
}

func main() {
	const currency = "COP"
	const minimumStock = 5

	products, err := loadProducts("products.json")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	products, err = inventory.AddProduct(
		products,
		domain.Product{
			SKU:   "KHAM001",
			Name:  "Khamra",
			Price: 200000,
			Stock: 3,
		},
	)

	if err != nil {
		fmt.Println("Could not add product:", err)
	}

	index, err := inventory.FindProductIndexBySKU(products, "CLUB001")

	if err != nil {
		fmt.Println("Product not found:", err)
	} else {
		fmt.Println("Product found")
		fmt.Println("SKU:", products[index].SKU)
		fmt.Println("Name:", products[index].Name)
		fmt.Println("Price:", products[index].Price)
		fmt.Println("Stock:", products[index].Stock)
	}

	fmt.Println()
	fmt.Println("PerfumeOps Catalog")
	fmt.Println("--------------------")

	totalInventoryValue := 0

	for index, product := range products {
		displayProduct(index, product, minimumStock)

		totalInventoryValue += product.Price * product.Stock
	}

	fmt.Println("--------------------")
	fmt.Printf("Total inventory value: %d %s\n", totalInventoryValue, currency)

	// Primera venta

	saleSKU := "YARA001"
	quantity := 1

	fmt.Println()
	fmt.Println("Processing sale...")
	fmt.Println("--------------------")

	products, revenue, err, productIndex := inventory.ProcessSaleBySKU(products, saleSKU, quantity)

	printSale(products, quantity, productIndex, err, revenue)

	// Segunda venta

	saleSKU2 := "KHAM001"
	quantity2 := 1

	fmt.Println()
	fmt.Println("Processing second sale...")
	fmt.Println("--------------------")

	products, revenue, err, productIndex2 := inventory.ProcessSaleBySKU(products, saleSKU2, quantity2)

	printSale(products, quantity2, productIndex2, err, revenue)

	err = saveProducts(products, "products.json")

	if err != nil {
		fmt.Println("Could not save products:", err)
		return
	}

	fmt.Println()
	fmt.Println("Products saved successfully")
}
