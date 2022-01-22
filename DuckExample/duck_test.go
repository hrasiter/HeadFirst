package duckexample

import "testing"

func Test_Dukc(t *testing.T) {
	t.Run("Test MallardDuck", func(t *testing.T) {
		md := MallardDuck{}
		fly := mallard.PerformFly()
		quack := mallard.PerformQuack()

		if fly != "Fly with wings" {
			t.Errorf("want: %s, got: %s", "Fly with wings", fly)
		}

		if quack != "Quack" {
			t.Errorf("want: %s, got: %s", "Quack", quack)
		}
	})
}
