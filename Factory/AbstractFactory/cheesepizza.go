package abstractfactory

import (
	"fmt"

	ingredients "github.com/hrasiter/HeadFirst/Factory/AbstractFactory/Ingredients"
)

type CheesePizza struct {
	factory     PizzaIngredientsFactory
	description string
	dough       ingredients.Dough
	cheese      ingredients.Cheese
	sauce       ingredients.Sauce
}

func (p *CheesePizza) GetDescription() string {
	return p.description
}

func (p *CheesePizza) prepare() {
	fmt.Println("Preparing ", p.description)
	p.dough = p.factory.createDough()
	p.cheese = p.factory.createCheese()
	p.sauce = p.factory.createSauce()
}

func (p *CheesePizza) bake() {

}

func (p *CheesePizza) cut() {

}

func (p *CheesePizza) box() {

}

func (p *CheesePizza) GetDough() ingredients.Dough {
	return p.dough
}

func (p *CheesePizza) GetCheese() ingredients.Cheese {
	return p.cheese
}

func (p *CheesePizza) GetSauce() ingredients.Sauce {
	return p.sauce
}
