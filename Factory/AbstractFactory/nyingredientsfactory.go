package abstractfactory

import ingredients "github.com/hrasiter/HeadFirst/Factory/AbstractFactory/Ingredients"

type NYIngredientsFactory struct {
}

func (f *NYIngredientsFactory) createDough() ingredients.Dough {
	return &ingredients.ThinCrustDough{}
}

func (f *NYIngredientsFactory) createSauce() ingredients.Sauce {
	return &ingredients.MarinaraSauce{}
}

func (f *NYIngredientsFactory) createCheese() ingredients.Cheese {
	return &ingredients.MozerallaCheese{}
}
