package abstractfactory

type NYPizzaStore struct {
	*PizzaStore
}

func NewNYPizzaStore() PizzaStoreInterface {
	store := &PizzaStore{}
	newstore := &NYPizzaStore{store}
	store.PizzaStoreInterface = newstore
	return newstore
}

func (s *NYPizzaStore) CreatePizza(t string) Pizza {
	pf := &NYIngredientsFactory{}
	var pizza Pizza
	if t == "Cheese" {
		pizza = &CheesePizza{
			description: "Newyork, Cheese",
			factory:     pf,
		}
	}
	return pizza
}
