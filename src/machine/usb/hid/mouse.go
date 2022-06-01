package hid

var Mouse *mouse

type mouse struct {
	buf *RingBuffer
}

func newMouse() *mouse {
	return &mouse{
		buf: NewRingBuffer(),
	}
}

func (m *mouse) callback() bool {
	if b, ok := m.buf.Get(); ok {
		sendUSBPacket(b[:5])
		return true
	}
	return false
}

// Move is a function that moves the mouse cursor.
func (m *mouse) Move(vx, vy int) {
	if vx == 0 && vy == 0 {
		return
	}

	if vx < -128 {
		vx = -128
	}
	if vx > 127 {
		vx = 127
	}

	if vy < -128 {
		vy = -128
	}
	if vy > 127 {
		vy = 127
	}

	m.buf.Put([]byte{
		0x01, 0x00, byte(vx), byte(vy), 0x00,
	})
}

// WHEEL controls the mouse wheel.
func (m *mouse) Wheel(v int) {
	if v == 0 {
		return
	}

	if v < -128 {
		v = -128
	}
	if v > 127 {
		v = 127
	}

	m.buf.Put([]byte{
		0x01, 0x00, 0x00, 0x00, byte(v),
	})
}

// WheelDown turns the mouse wheel down.
func (m *mouse) WheelDown() {
	m.Wheel(-1)
}

// WheelUp turns the mouse wheel up.
func (m *mouse) WheelUp() {
	m.Wheel(1)
}
