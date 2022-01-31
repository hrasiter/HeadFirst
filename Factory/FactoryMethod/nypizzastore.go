package factorymethod

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
	return &NYCheesePizza{
		description: "Newyork, Cheese",
	}
}
