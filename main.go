package main

import (
	"github.com/brutella/hc/log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	bla "github.com/rumblefrog/light-controller/accessory"
	"github.com/rumblefrog/light-controller/light"
	rpio "github.com/stianeikeland/go-rpio"
)

func main() {

	err := rpio.Open()

	if err != nil {
		log.Info.Panic("Unable to open RPIO pins")
	}

	defer rpio.Close()

	light.R = rpio.Pin(27)
	light.R.Mode(rpio.Output)

	light.G = rpio.Pin(17)
	light.G.Mode(rpio.Output)

	light.B = rpio.Pin(22)
	light.B.Mode(rpio.Output)

	light.R.WriteC(0)
	light.G.WriteC(0)
	light.B.WriteC(0)

	log.Debug.Enable()

	acc := bla.NewBedLight(accessory.Info{
		Name:         "Bed Light",
		SerialNumber: "ZBed1",
		Manufacturer: "Z",
		Model:        "ZBed",
	})

	acc.RGBLight.Green.OnValueRemoteUpdate(func(v int) {
		log.Info.Print(v)
	})

	config := hc.Config{
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
