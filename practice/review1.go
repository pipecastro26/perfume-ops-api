package main

import "fmt"

func ganancia_por_unidad(priceBuy, PriceSale int) int {
	return PriceSale - priceBuy
}
func ingresos(unitSale, PriceSale int) int {
	return PriceSale * unitSale
}
func stock_restantes(stock, unitSale int) int {
	return stock - unitSale
}
func stock_bajo(stock, unitSale, stockMinimun int) string {
	if stock_restantes(stock, unitSale) >= stockMinimun {
		return "Sin Riego"
	} else {
		return "Con Riesgo"
	}
}
func main() {
	product := "Fakhar"
	priceBuy := 120000
	PriceSale := 220000
	stock := 9
	unitSale := 3
	stockMinimun := 4

	fmt.Println("Producto: ", product)
	fmt.Println("Remaining stock: ", stock_restantes(stock, unitSale))
	fmt.Println("Revenue: ", ingresos(unitSale, PriceSale))
	fmt.Println("Profit per unit: ", ganancia_por_unidad(priceBuy, PriceSale))
	fmt.Println("Low stock: ", stock_bajo(stock, unitSale, stockMinimun))

}
