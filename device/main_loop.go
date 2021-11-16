package device

import (
	"limao/device/lifecycle"
	"limao/graphics"
)

func mainLoop(de *device) {

	go de.update()
	// loop ...
	for ev := range de.callEvents() {
		switch ev := ev.(type) {
		case lifecycle.Event:
			switch ev.Crosses(lifecycle.StageVisible) {
			case lifecycle.CrossOn:
				graphics.Load(de.ctx)
				de.app.OnStart()
				de.start.Broadcast()
				de.sendEvents(draw{})
			case lifecycle.CrossOff:
				graphics.Release()
				de.app.OnStop()
			}
		case reseizeEvent:
			graphics.SetViewPort(int(de.win.Width), int(de.win.Height))
		case draw:
			if de.ctx == nil || ev.External {
				continue
			}
			de.mu.Lock()
			graphics.Clear()
			de.app.Draw()
			graphics.Draw()
			de.drawDone <- struct{}{}
			de.sendEvents(draw{})
			de.mu.Unlock()
		}
	}
}
