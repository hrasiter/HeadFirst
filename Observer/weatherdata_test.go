package observer

import (
	"strconv"
	"testing"
)

type MockObserver struct {
	temp float64
	hum  float64
	pres float64
}

func (o *MockObserver) Update(temp float64, humidity float64, press float64) {
	o.temp = temp
	o.hum = humidity
	o.pres = press
}

func (o *MockObserver) Display() string {
	return strconv.FormatFloat(o.temp, 'f', -1, 64)
}

func Test_WeatherData(t *testing.T) {
	w := NewWeatherData()

	o1 := &MockObserver{}
	w.RegisterObserver(o1)

	if len(w.observers) != 1 {
		t.Errorf("Register does not work!")
	}

	o2 := &MockObserver{}
	w.RegisterObserver(o2)

	if len(w.observers) != 2 {
		t.Errorf("Register does not work!")
	}

	w.RemoveObserver(o1)

	if len(w.observers) != 1 {
		t.Errorf("Remove does not work!")
	}

	w.RemoveObserver(o2)

	if len(w.observers) != 0 {
		t.Errorf("Remove does not work!")
	}
}
