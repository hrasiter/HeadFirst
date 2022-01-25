package observer

import (
	"strconv"
)

type StatisticDisplay struct {
	subject Subject
	temp    float64
}

func NewStatisticDisplay(s Subject) *StatisticDisplay {
	o := &StatisticDisplay{subject: s}
	s.RegisterObserver(o)
	return o
}

func (d *StatisticDisplay) Display() string {
	return strconv.FormatFloat(d.temp, 'f', 1, 64)
}

func (d *StatisticDisplay) Update(temp float64, humidity float64, press float64) {
	d.temp = temp
}
