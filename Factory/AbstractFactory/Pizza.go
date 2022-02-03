package abstractfactory

import ingredients "github.com/hrasiter/HeadFirst/Factory/AbstractFactory/Ingredients"

type Pizza interface {
	GetDescription() string
	GetDough() ingredients.Dough
	GetSauce() ingredients.Sauce
	GetCheese() ingredients.Cheese
	prepare()
	bake()
	cut()
	box()
}
