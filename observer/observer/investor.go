package observer

import "fmt"

// ConcreteObserver
type Investor struct {
	Name string
}

func (i *Investor) Update(stockSymbol string, price float64) {
	fmt.Printf("Investor %s notified: Stock %s is now $%.2f\n", i.Name, stockSymbol, price)
}
