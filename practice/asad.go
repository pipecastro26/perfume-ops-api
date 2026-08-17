package main

import "fmt"

func main() {
	product := "Club de Nuit"
	Stock_inicial := 12
	Unidades_vendidas := 5
	Stock_mínimo := 6
	precioCompra := 130000
	precioVenta := 220000

	fmt.Println(product)
	fmt.Println(Stock_inicial - Unidades_vendidas)
	fmt.Println(Unidades_vendidas * precioVenta)
	fmt.Println(precioVenta - precioCompra)
	fmt.Println((precioVenta - precioCompra) * Unidades_vendidas)
	if Stock_inicial-Unidades_vendidas <= Stock_mínimo {
		fmt.Println("Low stock: true")
	} else {
		fmt.Println("Low stock: false")
	}

}
