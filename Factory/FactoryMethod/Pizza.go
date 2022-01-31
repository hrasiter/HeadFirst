package factorymethod

type Pizza interface {
	GetDescription() string
	prepare()
	bake()
	cut()
	box()
}
