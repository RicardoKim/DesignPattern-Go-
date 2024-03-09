package storage

import "github.com/RicardoKim/Study/DesignPattern/AbstractFactory/product"

type StorageFactory interface {
	CreateApple() product.Product
	CreateBanana() product.Product
}
