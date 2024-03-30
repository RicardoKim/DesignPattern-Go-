package menu

import "fmt"

type Tea struct{}

func (t *Tea) BoilWater() {
	fmt.Println("Boiling water")
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (t *Tea) AddCondiments() {
	fmt.Println("These Method should not be called")
}

func (c *Tea) CustomerWantsCondiments() bool {
	return false
}

func (c *Tea) PrepareRecipe() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	if c.CustomerWantsCondiments() {
		c.AddCondiments()
	}
}
