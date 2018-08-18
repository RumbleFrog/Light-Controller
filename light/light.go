package light

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// PATH - Path for pi-blaster
const PATH = "/dev/pi-blaster"

// Pin - Struct
type Pin struct {
	Pin int
}

// Pins receivers
var (
	Red   = &Pin{Pin: 27}
	Green = &Pin{Pin: 17}
	Blue  = &Pin{Pin: 22}
)

// Write - Write to pin
func (p Pin) Write(v float64) {
	Query := fmt.Sprintf("%d=%f", p.Pin, v)

	Cmd := exec.Command("echo", Query)

	OutBin, err := os.Create(PATH)

	if err != nil {
		log.Panic(err)
	}

	defer OutBin.Close()

	Cmd.Stdout = OutBin

	if err = Cmd.Start(); err != nil {
		log.Panic(err)
	}

	if err = Cmd.Wait(); err != nil {
		log.Println(err)
	}
}

// WriteAll - Write to all colors
func WriteAll(r, g, b float64) {
	Query := fmt.Sprintf("%d=%f %d=%f %d=%f", Red.Pin, r, Green.Pin, g, Blue.Pin, b)

	Cmd := exec.Command("echo", Query)

	OutBin, err := os.Create(PATH)

	if err != nil {
		log.Panic(err)
	}

	defer OutBin.Close()

	Cmd.Stdout = OutBin

	if err = Cmd.Start(); err != nil {
		log.Panic(err)
	}

	if err = Cmd.Wait(); err != nil {
		log.Println(err)
	}
}
