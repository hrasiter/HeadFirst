package duckexample

type Duck struct {
	DuckInterface
	flying   FlyBehaivior
	quacking QuackBehaivor
}

func (d *Duck) PerformFly() string {
	return d.flying.Fly()
}

func (d *Duck) PerformQuack() string {
	return d.quacking.Quack()
}

func (d *Duck) Swim() string {
	return "All ducks can swim"
}
