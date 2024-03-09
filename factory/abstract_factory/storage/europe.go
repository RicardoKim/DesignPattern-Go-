package storage

import "github.com/RicardoKim/Study/DesignPattern/AbstractFactory/product"

type EuropeStorage struct{}

func (es *EuropeStorage) CreateApple() product.Product {
	return &product.EuropeApple{}
}

func (es *EuropeStorage) CreateBanana() product.Product {
	return &product.EuropeBanana{}
}
