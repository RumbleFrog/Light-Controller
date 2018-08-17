package accessory

import (
	"github.com/brutella/hc/accessory"
	"github.com/rumblefrog/light-controller/service"
)

// BedLight - BedLight Accessory
type BedLight struct {
	*accessory.Accessory
	BedLight *service.BedLight
}

// NewBedLight - Returns the BedLight accessory
func NewBedLight(info accessory.Info) *BedLight {
	acc := BedLight{}
	acc.Accessory = accessory.New(info, accessory.TypeOther)
	acc.BedLight = service.NewBedLight()

	acc.BedLight.Red.SetValue(255)
	acc.BedLight.Green.SetValue(255)
	acc.BedLight.Blue.SetValue(255)

	acc.AddService(acc.BedLight.Service)

	return &acc
}
