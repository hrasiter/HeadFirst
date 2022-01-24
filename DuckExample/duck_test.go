package duckexample

import "testing"

func Test_Dukc(t *testing.T) {
	t.Run("Test MallardDuck", func(t *testing.T) {
		md := NewMallardDuck(&FlyWithWings{}, &Quack{})
		fly := md.PerformFly()
		quack := md.PerformQuack()
		disp := md.Display()

		if disp != "I'm Mallard Duck!!" {
			t.Errorf("Cannot display himself!")
		}

		if fly != "I'am flying with wings!" {
			t.Errorf("want: %s, got: %s", "I'am flying with wings!", fly)
		}

		if quack != "I am Quacking!" {
			t.Errorf("want: %s, got: %s", "I am Quacking!", quack)
		}
	})
}
