package storage

import (
	"fmt"

	"github.com/RicardoKim/Study/DesignPattern/FactoryMethod/product"
)

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) GetProduct(name string) (product.Product, error) {
	switch name {
	case "Apple":
		return product.NewApple(), nil
	case "Banana":
		return product.NewBanana(), nil
	default:
		return nil, fmt.Errorf("unknown product")
	}
}
