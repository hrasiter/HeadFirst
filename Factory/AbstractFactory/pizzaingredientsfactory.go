package abstractfactory

import ingredients "github.com/hrasiter/HeadFirst/Factory/AbstractFactory/Ingredients"

type PizzaIngredientsFactory interface {
	createDough() ingredients.Dough
	createSauce() ingredients.Sauce
	createCheese() ingredients.Cheese
}
