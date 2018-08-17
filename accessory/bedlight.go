package accessory

import (
	"github.com/brutella/hc/accessory"
	"github.com/rumblefrog/light-controller/service"
)

// BedLight - BedLight Accessory
type BedLight struct {
	*accessory.Accessory
	RGBLight *service.RGBLight
}

// NewBedLight - Returns the BedLight accessory
func NewBedLight(info accessory.Info) *BedLight {
	acc := BedLight{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.RGBLight = service.NewRGBLight()

	acc.RGBLight.Red.SetValue(0)
	acc.RGBLight.Green.SetValue(0)
	acc.RGBLight.Blue.SetValue(0)

	acc.AddService(acc.RGBLight.Service)

	return &acc
}
