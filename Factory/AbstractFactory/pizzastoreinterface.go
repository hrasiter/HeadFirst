package abstractfactory

type PizzaStoreInterface interface {
	CreatePizza(t string) Pizza
	OrderPizza(t string) Pizza
}
