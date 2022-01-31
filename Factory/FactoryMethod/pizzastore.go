package factorymethod

type PizzaStore struct {
	PizzaStoreInterface
}

func (p *PizzaStore) OrderPizza(t string) Pizza {
	pizza := p.CreatePizza(t)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza
}

// func (s *PizzaStore) CreatePizza(t string) Pizza {
// 	fmt.Println("Errorrrrrr")
// 	return &NYCheesePizza{}
// }
