package abstractfactory

type CAPizzaStore struct {
	*PizzaStore
}

func NewCAPizzaStore() PizzaStoreInterface {
	store := &PizzaStore{}
	cstore := &CAPizzaStore{store}
	store.PizzaStoreInterface = cstore
	return cstore
}

func (s *CAPizzaStore) CreatePizza(t string) Pizza {
	pf := &CAIngredientsFactory{}
	var pizza Pizza
	if t == "Cheese" {
		pizza = &CheesePizza{
			description: "California, Cheese",
			factory:     pf,
		}
	}
	return pizza
}
