package characteristic

import "github.com/brutella/hc/characteristic"

// TypeBlue - The characteristic index I suppose
const TypeBlue = "680"

// Blue Export handle
type Blue struct {
	*characteristic.Int
}

// NewBlue - Create handle
func NewBlue() *Blue {
	char := characteristic.NewInt(TypeBlue)
	char.Format = characteristic.FormatUInt8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetMinValue(0)
	char.SetMaxValue(255)
	char.SetStepValue(1)
	char.SetValue(0)

	return &Blue{char}
}
