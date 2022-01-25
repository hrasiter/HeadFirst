package observer

type CurrentConditionDisplay struct {
	subject Subject
}

func (d *CurrentConditionDisplay) Display() string {
	return ""
}
