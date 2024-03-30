package beverage

import (
	"fmt"

	"github.com/RicardoKim/Study/DesignPattern/TemplateMethodPattern/menu"
)

type BeverageFactory struct{}

func (bf *BeverageFactory) OrderBeverage(beverageType string) Beverage {
	switch beverageType {
	case "coffee":
		return &menu.Coffee{}
	case "tea":
		return &menu.Tea{}
	default:
		fmt.Println("Sorry, we don't prepare", beverageType)
		return nil
	}
}
