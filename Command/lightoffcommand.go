package command

type LightOffCommand struct {
	light *Light
}

func (offcmd *LightOffCommand) execute() {
	offcmd.light.SetLightOn(false)
}
