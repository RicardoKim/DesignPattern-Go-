package beverage

type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
	CustomerWantsCondiments() bool
	PrepareRecipe()
}
