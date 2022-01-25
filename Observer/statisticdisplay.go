package observer

type StatisticDisplay struct {
	subject Subject
}

func (d *StatisticDisplay) Display() string {
	return ""
}
