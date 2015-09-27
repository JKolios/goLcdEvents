// +build ignore

package main

import (
	"flag"

	"github.com/Jkolios/goLcdEvents/Godeps/_workspace/src/github.com/kidoman/embd"

	_ "github.com/Jkolios/goLcdEvents/Godeps/_workspace/src/github.com/kidoman/embd/host/all"
)

func main() {
	flag.Parse()

	embd.InitGPIO()
	defer embd.CloseGPIO()

	embd.SetDirection(10, embd.Out)
	embd.DigitalWrite(10, embd.High)
}
