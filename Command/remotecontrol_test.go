package command

import "testing"

func Test_RemoteControl(t *testing.T) {
	t.Run("Testing Simple Command Remote Control", func(t *testing.T) {
		remote := &SimpleRemoteControl{}
		light := &Light{}
		state := light.GetState()

		if state != false {
			t.Errorf("Light should be Off!")
		}
		remote.SetCommand(&LightOnCommand{light})
		remote.ButtonPressed()
		state = light.GetState()

		if state != true {
			t.Errorf("Light should be On!")
		}

		remote.SetCommand(&LightOffCommand{light})
		remote.ButtonPressed()
		state = light.GetState()

		if state != false {
			t.Errorf("Light should be Off!")
		}

	})
}
