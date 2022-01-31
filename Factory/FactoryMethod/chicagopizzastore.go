package factorymethod

type ChicagoPizzaStore struct {
	*PizzaStore
}

func NewChicagoPizzaStore() PizzaStoreInterface {
	store := &PizzaStore{}
	cstore := &ChicagoPizzaStore{store}
	store.PizzaStoreInterface = cstore
	return cstore
}

func (s *ChicagoPizzaStore) CreatePizza(t string) Pizza {
	return &ChicagoCheesePizza{
		description: "Chicago, Cheese",
	}
}
