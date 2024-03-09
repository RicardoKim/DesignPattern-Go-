package main

import (
	"fmt"

	"github.com/RicardoKim/Study/DesignPattern/AbstractFactory/storage"
)

func main() {
	europeStorage := &storage.EuropeStorage{}
	southeastAsiaStorage := &storage.SouthEastAsiaStorage{}

	// 유럽 애플 생성
	europeApple := europeStorage.CreateApple()
	fmt.Printf("Product: %s, Price: $%.2f\n", europeApple.Name(), europeApple.Price())

	// 동남아 바나나 생성
	southeastAsiaBanana := southeastAsiaStorage.CreateBanana()
	fmt.Printf("Product: %s, Price: $%.2f\n", southeastAsiaBanana.Name(), southeastAsiaBanana.Price())
}
