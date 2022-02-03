package abstractfactory

import (
	"testing"

	ingredients "github.com/hrasiter/HeadFirst/Factory/AbstractFactory/Ingredients"
)

func Test_PizzaFactory(t *testing.T) {
	t.Run("Testing NyPizza", func(t *testing.T) {
		store := NewNYPizzaStore()
		pizza := store.OrderPizza("Cheese")

		if pizza.GetDescription() != "Newyork, Cheese" {
			t.Errorf("Cannot create Newyork Type Cheese")
		}

		d := pizza.GetDough()

		if _, ok := d.(*ingredients.ThinCrustDough); !ok {
			t.Errorf("NYPizza should have thin crust dough")
		}
	})

	t.Run("Testing ChicagoPizza", func(t *testing.T) {
		store := NewCAPizzaStore()
		pizza := store.OrderPizza("Cheese")

		if pizza.GetDescription() != "California, Cheese" {
			t.Errorf("Cannot create California Type Cheese")
		}

		d := pizza.GetDough()

		if _, ok := d.(*ingredients.ThickCrustDough); !ok {
			t.Errorf("CAPizza should have thick crust dough")
		}
	})
}
