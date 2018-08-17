package service

import (
	"github.com/brutella/hc/service"
	"github.com/rumblefrog/light-controller/characteristic"
)

// TypeRGBLight - Index?
const TypeRGBLight = "C89"

// RGBLight - Info related to operating the light
type RGBLight struct {
	*service.Service

	Red   *characteristic.Red
	Green *characteristic.Green
	Blue  *characteristic.Blue
}

// NewRGBLight - Creates a new handle
func NewRGBLight() *RGBLight {
	svc := RGBLight{}
	svc.Service = service.New(TypeRGBLight)

	svc.Red = characteristic.NewRed()
	svc.AddCharacteristic(svc.Red.Characteristic)

	svc.Green = characteristic.NewGreen()
	svc.AddCharacteristic(svc.Green.Characteristic)

	svc.Blue = characteristic.NewBlue()
	svc.AddCharacteristic(svc.Blue.Characteristic)

	return &svc
}
