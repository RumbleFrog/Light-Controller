package light

import rpio "github.com/stianeikeland/go-rpio"

// RGB Pins
var (
	R rpio.Pin
	G rpio.Pin
	B rpio.Pin
)

// WriteC - Sets a pin's color
func (p rpio.Pin) WriteC(v rpio.State) {
	p.Write(v)

	// Broadcast to whatever needs a state update
}
