package storage

import "github.com/RicardoKim/Study/DesignPattern/AbstractFactory/product"

type SouthEastAsiaStorage struct{}

func (ss *SouthEastAsiaStorage) CreateApple() product.Product {
	return &product.SoutheastAsiaApple{}
}

func (ss *SouthEastAsiaStorage) CreateBanana() product.Product {
	return &product.SoutheastAsiaBanana{}
}
