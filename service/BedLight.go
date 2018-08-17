package service

import (
	"github.com/brutella/hc/service"
	"github.com/rumblefrog/light-controller/characteristic"
)

// TypeBedLight - Index?
const TypeBedLight = "C89"

// BedLight - Info related to operating the light
type BedLight struct {
	*service.Service

	Red   *characteristic.Red
	Green *characteristic.Green
	Blue  *characteristic.Blue
}

// NewBedLight - Creates a new handle
func NewBedLight() *BedLight {
	svc := BedLight{}
	svc.Service = service.New(TypeBedLight)

	svc.Red = characteristic.NewRed()
	svc.AddCharacteristic(svc.Red.Characteristic)

	svc.Green = characteristic.NewGreen()
	svc.AddCharacteristic(svc.Green.Characteristic)

	svc.Blue = characteristic.NewBlue()
	svc.AddCharacteristic(svc.Blue.Characteristic)

	return &svc
}
