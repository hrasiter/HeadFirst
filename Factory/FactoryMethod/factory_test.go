package factorymethod

import "testing"

func Test_PizzaFactory(t *testing.T) {
	t.Run("Testing NyPizza", func(t *testing.T) {
		store := NewNYPizzaStore()
		pizza := store.OrderPizza("Cheese")

		if pizza.GetDescription() != "Newyork, Cheese" {
			t.Errorf("Cannot create Newyork Type Cheese")
		}
	})

	t.Run("Testing ChicagoPizza", func(t *testing.T) {
		store := NewChicagoPizzaStore()
		pizza := store.OrderPizza("Cheese")

		if pizza.GetDescription() != "Chicago, Cheese" {
			t.Errorf("Cannot create Chicago Type Cheese")
		}
	})
}
