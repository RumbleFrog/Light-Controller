package characteristic

import (
	"github.com/brutella/hc/characteristic"
)

// TypeGreen - The characteristic index I suppose
const TypeGreen = "679"

// Green Export handle
type Green struct {
	*characteristic.Int
}

// NewGreen - Create handle
func NewGreen() *Green {
	char := characteristic.NewInt(TypeGreen)
	char.Format = characteristic.FormatUInt8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetMinValue(0)
	char.SetMaxValue(255)
	char.SetStepValue(1)
	char.SetValue(0)

	return &Green{char}
}
