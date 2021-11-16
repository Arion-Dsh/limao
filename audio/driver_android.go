//go:build android
// +build android

package audio

import (
	"limao/audio/oboe"
	"sync"
)

func newDriver(rate, depth, chs int) driver {
	o := &driverTmp{
		cond:      sync.NewCond(new(sync.Mutex)),
		buff:      []float32{},
		rate:      rate,
		depth:     depth,
		chs:       chs,
		framesNum: 128,
		device:    oboe.New(),
	}
	return o
}
