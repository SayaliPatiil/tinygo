package main

import (
	"machine"
	"machine/usb/hid"
	"time"
)

func main() {
	button := machine.BUTTON
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	kb := hid.Keyboard

	for {
		if !button.Get() {
			kb.Write([]byte("tinygo"))
			time.Sleep(200 * time.Millisecond)
		}
	}
}
