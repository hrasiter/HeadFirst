package observer

import (
	"strconv"
)

type CurrentConditionDisplay struct {
	subject Subject
	temp    float64
}

func NewCurrentConditionDisplay(s Subject) *CurrentConditionDisplay {
	o := &CurrentConditionDisplay{subject: s}
	s.RegisterObserver(o)
	return o
}

func (d *CurrentConditionDisplay) Display() string {
	return strconv.FormatFloat(d.temp, 'f', 1, 64)
}

func (d *CurrentConditionDisplay) Update(temp float64, humidity float64, press float64) {
	d.temp = temp
}
