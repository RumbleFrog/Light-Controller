package main

import (
	"log"

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
		log.Fatal("Unable to open RPIO pins")
	}

	defer rpio.Close()

	R = rpio.Pin(27)
	R.Mode(rpio.Output)

	G = rpio.Pin(17)
	G.Mode(rpio.Output)

	B = rpio.Pin(22)
	B.Mode(rpio.Output)

	acc := bla.NewBedLight(accessory.Info{
		Name:         "Bed Light",
		Manufacturer: "Z",
	})

	config := hc.Config{
		Pin: "00011019",
	}

	t, err := hc.NewIPTransport(config, acc.Accessory)

	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
