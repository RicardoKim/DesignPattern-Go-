package main

import (
	"github.com/RicardoKim/Study/DesignPattern/Observer/observer"
	"github.com/RicardoKim/Study/DesignPattern/Observer/subject"
)

func main() {
	ticker := &subject.StockTicker{}
	ticker.SetStockSymbol("Samsung")

	investor1 := &observer.Investor{Name: "John"}
	investor2 := &observer.Investor{Name: "Doe"}

	ticker.RegisterObserver(investor1)
	ticker.RegisterObserver(investor2)

	ticker.SetPrice(1500.0)
	ticker.SetPrice(1550.0)
}
