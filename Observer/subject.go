package observer

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	Notify()
}
