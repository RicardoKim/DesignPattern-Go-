package product

type EuropeBanana struct{}

func (eb *EuropeBanana) Name() string {
	return "Europe Banana"
}

func (eb *EuropeBanana) Price() float64 {
	return 2.50
}

type SoutheastAsiaBanana struct{}

func (sb *SoutheastAsiaBanana) Name() string {
	return "Southeast Asis Banana"
}

func (sb *SoutheastAsiaBanana) Price() float64 {
	return 1.50
}
