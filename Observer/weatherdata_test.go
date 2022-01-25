package observer

import (
	"testing"
)

type MockObserver struct {
}

func (o *MockObserver) Update(float64, float64, float64) {

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
