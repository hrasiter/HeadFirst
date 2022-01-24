package duckexample

type Quack struct {
}

func (q *Quack) Quack() string {
	return "I am Quacking!"
}
