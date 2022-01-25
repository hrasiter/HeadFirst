package observer

import "testing"

func Test_Observer(t *testing.T) {
	t.Run("test publish subsribe", func(t *testing.T) {

		s := NewWeatherData()
		o1 := NewCurrentConditionDisplay(s)
		o2 := NewStatisticDisplay(s)

		s.SetTemp(10.0)

		s.MeasurementsChanged()

		r1 := o1.Display()
		r2 := o2.Display()

		if r1 != "10.0" {
			t.Errorf("observer1 is not being notifed! want: %s, got: %s", "10.0", r1)
		}

		if r2 != "10.0" {
			t.Errorf("observer2 is not being notifed! want: %s, got: %s", "10.0", r2)
		}
	})
}
