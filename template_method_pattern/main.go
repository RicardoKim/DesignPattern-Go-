package main

import (
	"fmt"

	"github.com/RicardoKim/Study/DesignPattern/TemplateMethodPattern/beverage"
)

func main() {
	factory := &beverage.BeverageFactory{}

	beverageType := "coffee"
	beverage := factory.OrderBeverage(beverageType)
	if beverage != nil {
		fmt.Printf("Preparing %s...\n", beverageType)
		beverage.PrepareRecipe()
	}

	beverageType = "tea"
	beverage = factory.OrderBeverage(beverageType)
	if beverage != nil {
		fmt.Printf("Preparing %s...\n", beverageType)
		beverage.PrepareRecipe()
	}
}
