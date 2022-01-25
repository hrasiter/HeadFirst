package observer

type WeatherData struct {
	observers []Observer
	temp      float64
	humidity  float64
	pressure  float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make([]Observer, 0),
	}
}

func (w *WeatherData) Notify() {
	for _, d := range w.observers {
		d.Update(w.temp, w.humidity, w.pressure)
	}
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	// i := sort.Search(len(w.observers), func(i int) bool {
	// 	return w.observers[i] == o
	// })

	for i, _ := range w.observers {
		if w.observers[i] == o {
			w.remove(i)
			return
		}
	}
}

func (w *WeatherData) SetTemp(temp float64) {
	w.temp = temp
}

func (w *WeatherData) MeasurementsChanged() {
	w.Notify()
}

func (w *WeatherData) remove(i int) {
	copy(w.observers[i:], w.observers[i+1:])
	w.observers = append(w.observers[:i], w.observers[i+1:]...)

}
