// Package all conviniently loads all the inbuilt/supported host drivers.
package all

import (
	_ "github.com/JKolios/goLcdEvents/Godeps/_workspace/src/github.com/kidoman/embd/host/bbb"
	_ "github.com/JKolios/goLcdEvents/Godeps/_workspace/src/github.com/kidoman/embd/host/rpi"
)