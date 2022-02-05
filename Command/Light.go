package command

type LightState int

const (
	LightOn LightState = iota
	LightOff
)

type Light struct {
	isLightOn bool
}

func (l *Light) GetState() bool {
	return l.isLightOn
}

func (l *Light) SetLightOn(isOn bool) {
	l.isLightOn = isOn
}
