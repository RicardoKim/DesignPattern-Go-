package menu

import "fmt"

type Coffee struct{}

// PourInCpu implements beverage.Beverage.
func (*Coffee) PourInCpu() {
	panic("unimplemented")
}

func (c *Coffee) BoilWater() {
	fmt.Println("Boiling water")
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (c *Coffee) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}

func (t *Coffee) CustomerWantsCondiments() bool {
	return true // 사용자가 항상 조미료를 원한다고 가정
}

// PrepareRecipe 메서드 수정
func (t *Coffee) PrepareRecipe() {
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	if t.CustomerWantsCondiments() {
		t.AddCondiments()
	}
}
