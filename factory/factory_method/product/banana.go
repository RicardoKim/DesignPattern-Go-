package product

type Banana struct {
	productName  string
	productPrice float64
}

func NewBanana() *Banana {
	return &Banana{"Banana", 1.99}
}

func (b *Banana) Name() string {
	return b.productName
}

func (b *Banana) Price() float64 {
	return b.productPrice
}
