package characteristic

import (
	"github.com/brutella/hc/characteristic"
)

// TypeRed - The characteristic index I suppose
const TypeRed = "678"

// Red Export handle
type Red struct {
	*characteristic.Int
}

// NewRed - Create handle
func NewRed() *Red {
	char := characteristic.NewInt(TypeRed)
	char.Format = characteristic.FormatUInt8

	char.SetMinValue(0)
	char.SetMaxValue(255)
	char.SetStepValue(1)
	char.SetValue(0)

	return &Red{char}
}
