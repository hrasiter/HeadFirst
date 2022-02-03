package abstractfactory

import ingredients "github.com/hrasiter/HeadFirst/Factory/AbstractFactory/Ingredients"

type CAIngredientsFactory struct {
}

func (f *CAIngredientsFactory) createDough() ingredients.Dough {
	return &ingredients.ThickCrustDough{}
}

func (f *CAIngredientsFactory) createSauce() ingredients.Sauce {
	return &ingredients.TomatoSauce{}
}

func (f *CAIngredientsFactory) createCheese() ingredients.Cheese {
	return &ingredients.ReggianoCheese{}
}
