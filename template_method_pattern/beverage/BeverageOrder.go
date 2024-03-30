package beverage

type BeverageOrder interface {
	OrderVeverage(beverageType string) Beverage
}
