package device

import (
	"limao/device/lifecycle"
	"limao/graphics/gl"
	"limao/window"
	"sync"
	"time"
)

var theDevice *device

// Run ...
func Run(app appI, opts *window.Opts) {
	d := &device{
		start:      sync.NewCond(new(sync.Mutex)),
		drawDone:   make(chan struct{}),
		fireEvents: make(chan interface{}),

		app: app,
		win: &win{opts},

		mu: sync.Mutex{},
	}
	d.events = pump(d.fireEvents)

	theDevice = d
	theDevice.run()
}

type device struct {
	start          *sync.Cond
	ctx            gl.Context
	events         chan interface{}
	fireEvents     chan interface{}
	lifecycleStage lifecycle.Stage

	drawDone chan struct{}
	win      *win

	resize  bool
	suspend bool

	mu sync.Mutex

	// app ...
	app appI
}

type appI interface {
	OnStart()
	Update()
	Draw()
	OnStop()
}

func (de *device) update() {

	de.start.L.Lock()
	de.start.Wait()
	de.start.L.Unlock()

	ticker := time.NewTicker(time.Second / 60)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			de.mu.Lock()
			de.app.Update()
			de.mu.Unlock()
		}
	}
}

func (de *device) sendLifecycle(to lifecycle.Stage) {
	if de.lifecycleStage == to {
		return
	}
	de.events <- lifecycle.Event{
		From: de.lifecycleStage,
		To:   to,
	}
	de.lifecycleStage = to
}

func (de *device) sendEvents(e interface{}) {
	de.events <- e
}

func (de *device) callEvents() <-chan interface{} {
	return de.fireEvents
}

type stopPumping struct{}

// pump returns a channel src such that sending on src will eventually send on
// dst, in order, but that src will always be ready to send/receive soon, even
// if dst currently isn't. It is effectively an infinitely buffered channel.
//
// In particular, goroutine A sending on src will not deadlock even if goroutine
// B that's responsible for receiving on dst is currently blocked trying to
// send to A on a separate channel.
//
// Send a stopPumping on the src channel to close the dst channel after all queued
// events are sent on dst. After that, other goroutines can still send to src,
// so that such sends won't block forever, but such events will be ignored.
func pump(dst chan interface{}) (src chan interface{}) {
	src = make(chan interface{})
	go func() {
		// initialSize is the initial size of the circular buffer. It must be a
		// power of 2.
		const initialSize = 16
		i, j, buf, mask := 0, 0, make([]interface{}, initialSize), initialSize-1

		srcActive := true
		for {
			maybeDst := dst
			if i == j {
				maybeDst = nil

			}
			if maybeDst == nil && !srcActive {
				// Pump is stopped and empty.
				break

			}

			select {
			case maybeDst <- buf[i&mask]:
				buf[i&mask] = nil
				i++

			case e := <-src:
				if _, ok := e.(stopPumping); ok {
					srcActive = false
					continue
				}
				if !srcActive {
					continue
				}

				// Allocate a bigger buffer if necessary.
				if i+len(buf) == j {
					b := make([]interface{}, 2*len(buf))
					n := copy(b, buf[j&mask:])
					copy(b[n:], buf[:j&mask])
					i, j = 0, len(buf)
					buf, mask = b, len(b)-1
				}
				buf[j&mask] = e
				j++
			}
		}

		close(dst)
		// Block forever.
		for range src {
		}

	}()
	return src

}
