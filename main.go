package main

import (
	"github.com/brutella/hc/log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	bla "github.com/rumblefrog/light-controller/accessory"
	rpio "github.com/stianeikeland/go-rpio"
)

// RGB Pins
var (
	R rpio.Pin
	G rpio.Pin
	B rpio.Pin
)

func main() {

	err := rpio.Open()

	if err != nil {
		log.Info.Panic("Unable to open RPIO pins")
	}

	defer rpio.Close()

	R = rpio.Pin(27)
	R.Mode(rpio.Output)

	G = rpio.Pin(17)
	G.Mode(rpio.Output)

	B = rpio.Pin(22)
	B.Mode(rpio.Output)

	R.Write(0)
	G.Write(0)
	B.Write(0)

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
