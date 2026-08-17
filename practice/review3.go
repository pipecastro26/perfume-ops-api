package main

import "fmt"

func main() {
	const stockMinimun = 5
	products := []string{
		"Yara",
		"Khamrah",
		"Asad",
		"Fakhar",
	}
	stocks := []int{
		8,
		2,
		0,
		6,
	}
	for index, product := range products {
		if stocks[index] >= stockMinimun {
			fmt.Println(index+1, ". ", product, " - Available")
		} else if stocks[index] < stockMinimun && stocks[index] > 0 {
			fmt.Println(index+1, ". ", product, " - Low stock")
		} else {
			fmt.Println(index+1, ". ", product, " - Out of stock")
		}
	}
}
