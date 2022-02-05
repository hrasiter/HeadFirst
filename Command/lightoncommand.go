package command

type LightOnCommand struct {
	light *Light
}

func (lc *LightOnCommand) execute() {
	lc.light.SetLightOn(true)
}
