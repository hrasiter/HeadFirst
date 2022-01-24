package duckexample

type MallardDuck struct {
	Duck
}

func NewMallardDuck(fb FlyBehaivior, qb QuackBehaivor) *MallardDuck {
	return &MallardDuck{
		Duck{
			flying:   fb,
			quacking: qb,
		},
	}
}

func (md *MallardDuck) Display() string {
	return "I'm Mallard Duck!!"
}
