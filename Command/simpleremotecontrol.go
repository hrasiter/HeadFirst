package command

type SimpleRemoteControl struct {
	command Command
}

func (rc *SimpleRemoteControl) SetCommand(c Command) {
	rc.command = c
}

func (rc *SimpleRemoteControl) ButtonPressed() {
	rc.command.execute()
}
