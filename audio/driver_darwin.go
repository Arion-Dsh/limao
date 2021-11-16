//go:build darwin
// +build darwin

package audio

import (
	"limao/audio/avaudio"
	"sync"
)

func newDriver(rate, depth, chs int) driver {
	d := &driverTmp{
		cond:      sync.NewCond(new(sync.Mutex)),
		buff:      []float32{},
		rate:      rate,
		depth:     depth,
		chs:       chs,
		framesNum: 4096,
		device:    avaudio.New(),
	}
	return d
}
