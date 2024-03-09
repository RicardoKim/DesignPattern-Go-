package product

type EuropeApple struct{}

func (ea *EuropeApple) Name() string {
	return "Europe Apple"
}

func (ea *EuropeApple) Price() float64 {
	return 2.50
}

type SoutheastAsiaApple struct{}

func (sa *SoutheastAsiaApple) Name() string {
	return "Southeast Asis Apple"
}

func (sa *SoutheastAsiaApple) Price() float64 {
	return 1.50
}
