package observer

// Observer interface
type IObserver interface {
	Update(stockSymbol string, price float64)
}
