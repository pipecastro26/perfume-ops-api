package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func addProduct(products []Product, product Product) ([]Product, error) {
	for _, existingProduct := range products {
		if existingProduct.SKU == product.SKU {
			return products, errors.New("product already exists")
		}
	}

	return append(products, product), nil
}

func displayProduct(index int, product Product, minimumStock int) {
	fmt.Printf(
		"%d. %s - %s - %d COP - Stock: %d",
		index+1,
		product.SKU,
		product.Name,
		product.Price,
		product.Stock,
	)

	if product.Stock <= minimumStock {
		fmt.Print(" - WARNING: Low stock")
	}

	fmt.Println()
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
func processSaleBySKU(
	products []Product,
	sku string,
	quantity int,
) ([]Product, int, error, int) {
	index, err := findProductIndexBySKU(products, sku)
	if err != nil {
		return products, 0, errors.New("product not found"), 0
	}
	err2 := products[index].DecreaseStock(quantity)
	if err2 != nil {
		return products, 0, errors.New("quantity must be greater than zero"), 0
	}
	return products, products[index].Revenue(quantity), nil, index

}
func (p Product) Revenue(quantity int) int {
	return p.Price * quantity
}
func (p *Product) IncreaseStock(quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	p.Stock += quantity
	return nil
}
func (p *Product) DecreaseStock(quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	if quantity > p.Stock {
		return errors.New("insufficient stock")
	}

	p.Stock -= quantity
	return nil
}
func (p *Product) ChangePrice(newPrice int) error {
	if newPrice <= 0 {
		return errors.New("price must be greater than zero")
	}

	p.Price = newPrice
	return nil
}
func findProductIndexBySKU(products []Product, sku string) (int, error) {
	for index, product := range products {
		if product.SKU == sku {
			return index, nil
		}
	}

	return 0, errors.New("product not found")
}
func main() {
	const currency = "COP"
	const minimumStock = 5

	products, err4 := loadProducts("products.json")
	if err4 != nil {
		fmt.Print("Error", err4)
		return
	}
	products, err5 := addProduct(products, Product{SKU: "KHAM001", Name: "Khamra", Price: 200000, Stock: 3})
	if err5 != nil {
		fmt.Println("Could not add product:", err5)
	}
	index, err := findProductIndexBySKU(products, "CLUB001")
	if err != nil {
		fmt.Print("No se encontro")
	} else {
		fmt.Println(products[index])
	}
	fmt.Println()
	fmt.Println("PerfumeOps Catalog")
	fmt.Println("--------------------")

	totalInventoryValue := 0

	for index, product := range products {
		displayProduct(
			index,
			product,
			minimumStock,
		)

		totalInventoryValue += product.Price * product.Stock
	}

	fmt.Println("--------------------")
	fmt.Printf(
		"Total inventory value: %d %s\n",
		totalInventoryValue,
		currency,
	)

	saleSku := "YARA001"
	quantity := 1000

	fmt.Println()
	fmt.Println("Processing sale...")
	fmt.Println("--------------------")
	products, revenue, err6, productIndex := processSaleBySKU(products, saleSku, quantity)
	if err6 != nil {
		fmt.Println("Sale could not be processed:", err6)
	} else {

		fmt.Println("Sale processed successfully")
		fmt.Println("Product:", products[productIndex].Name)
		fmt.Println("Units sold:", quantity)
		fmt.Println("Remaining stock:", products[productIndex].Stock)
		fmt.Println("Revenue:", revenue)
	}

	// Segunda venta
	saleSku2 := "KHAM001"
	quantity2 := 0

	fmt.Println()
	fmt.Println("Processing second sale...")
	fmt.Println("--------------------")
	products, revenue, err7, productIndex2 := processSaleBySKU(products, saleSku2, quantity2)
	if err7 != nil {
		fmt.Println("Sale could not be processed:", err7)
	} else {
		fmt.Println("Sale processed successfully")
		fmt.Println("Product:", products[productIndex2].Name)
		fmt.Println("Units sold:", quantity2)
		fmt.Println("Remaining stock:", products[productIndex2].Stock)
		fmt.Println("Revenue:", revenue)
	}
	err8 := saveProducts(products, "products.json")
	if err8 != nil {
		fmt.Print("No se puede guardar")
	}

}
