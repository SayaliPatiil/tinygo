package hid

import (
	"errors"
	"machine"
)

// from usb-hid-keyboard.go
var (
	ErrInvalidCodepoint = errors.New("invalid Unicode codepoint")
	ErrInvalidKeycode   = errors.New("invalid keyboard keycode")
	ErrInvalidUTF8      = errors.New("invalid UTF-8 encoding")
	ErrKeypressMaximum  = errors.New("maximum keypresses exceeded")
)

// from usb-hid.go
var (
	ErrHIDInvalidPort    = errors.New("invalid USB port")
	ErrHIDInvalidCore    = errors.New("invalid USB core")
	ErrHIDReportTransfer = errors.New("failed to transfer HID report")
)

const (
	hidEndpoint = 4
)

func init() {
	machine.EnableHID(callback)

	Keyboard = newKeyboard()
	Mouse = newMouse()
}

func callback() {
	if done := Keyboard.callback(); done {
		return
	}

	if done := Mouse.callback(); done {
		return
	}
}

func sendUSBPacket(b []byte) {
	machine.SendUSBHIDPacket(hidEndpoint, b)
}
