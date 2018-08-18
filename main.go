package main

import (
	"github.com/brutella/hc/log"
	colorful "github.com/lucasb-eyer/go-colorful"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/rumblefrog/light-controller/api"
	"github.com/rumblefrog/light-controller/light"
)

var acc *accessory.Lightbulb

func main() {
	go api.Register()

	acc = accessory.NewLightbulb(accessory.Info{
		Name:         "Bed Light",
		SerialNumber: "ZBed1",
		Manufacturer: "Z",
		Model:        "ZBed",
	})

	acc.Lightbulb.On.OnValueRemoteUpdate(func(v bool) {
		if v == false {
			light.WriteAll(0.0, 0.0, 0.0)
		} else {
			updateLight()
		}
	})

	acc.Lightbulb.Brightness.OnValueRemoteUpdate(func(v int) {
		updateLight()
	})

	acc.Lightbulb.Saturation.OnValueRemoteUpdate(func(v float64) {
		updateLight()
	})

	acc.Lightbulb.Hue.OnValueRemoteUpdate(func(v float64) {
		updateLight()
	})

	config := hc.Config{
		Port:        "38631",
		Pin:         "00011019",
		StoragePath: "./db",
	}

	t, err := hc.NewIPTransport(config, acc.Accessory)

	if err != nil {
		log.Info.Panic(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}

func updateLight() {
	R, G, B := colorful.Hsv(
		acc.Lightbulb.Hue.GetValue(),
		acc.Lightbulb.Saturation.GetValue()/100.0,
		float64(acc.Lightbulb.Brightness.GetValue())/100.0,
	).RGB255()

	light.WriteAll(
		float64(R)/255.0,
		float64(G)/255.0,
		float64(B)/255.0,
	)
}
