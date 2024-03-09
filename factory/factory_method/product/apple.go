package product

type Apple struct {
	productName  string
	productPrice float64
}

func NewApple() *Apple {
	return &Apple{"Apple", 1.99}
}

func (a *Apple) Name() string {
	return a.productName
}

func (a *Apple) Price() float64 {
	return a.productPrice
}
