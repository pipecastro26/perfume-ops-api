package domain

import (
	"errors"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
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
