package inventory

import (
	"errors"

	"github.com/pipecastro26/perfume-ops-api/internal/domain"
)

func AddProduct(products []domain.Product, product domain.Product) ([]domain.Product, error) {
	for _, existingProduct := range products {
		if existingProduct.SKU == product.SKU {
			return products, errors.New("product already exists")
		}
	}

	return append(products, product), nil
}
func FindProductIndexBySKU(products []domain.Product, sku string) (int, error) {
	for index, product := range products {
		if product.SKU == sku {
			return index, nil
		}
	}

	return 0, errors.New("product not found")
}
func ProcessSaleBySKU(products []domain.Product, sku string, quantity int) ([]domain.Product, int, error, int) {
	index, err := FindProductIndexBySKU(products, sku)

	if err != nil {
		return products, 0, err, 0
	}

	err = products[index].DecreaseStock(quantity)

	if err != nil {
		return products, 0, err, index
	}

	revenue := products[index].Revenue(quantity)

	return products, revenue, nil, index
}
