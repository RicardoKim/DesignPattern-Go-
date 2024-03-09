package main

import (
	"fmt"

	"github.com/RicardoKim/Study/DesignPattern/FactoryMethod/storage"
)

func main() {
	storage := storage.NewStorage()
	apple, err := storage.GetProduct("Apple")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Product: %s, Price: $%.2f\n", apple.Name(), apple.Price())

	banana, err := storage.GetProduct("Banana")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Product: %s, Price: $%.2f\n", banana.Name(), banana.Price())
}
