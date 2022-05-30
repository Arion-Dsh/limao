//go:build (linux && !android) || darwin || windows
// +build linux,!android darwin windows

package device

import (
	"limao/device/lifecycle"
	"limao/graphics/gl"
	"limao/input"
	"runtime"
	"sync"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var fwWin *glfw.Window
var monitor *glfw.Monitor
var videoMode *glfw.VidMode
var mu sync.Mutex

// hack for begin window size
var ow int = 4
var oh int = 4

func (de *device) run() {

	runtime.LockOSThread()

	glctx, worker := gl.NewContext()

	de.ctx = glctx

	err := glfw.Init()

	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Visible, glfw.False)
	fwWin, err = glfw.CreateWindow(ow, oh, de.win.Title, nil, nil)
	if err != nil {
		panic(err)
	}

	// make gle context()
	fwWin.MakeContextCurrent()
	fwWin.SetInputMode(glfw.StickyMouseButtonsMode, glfw.True)
	fwWin.SetInputMode(glfw.StickyKeysMode, glfw.True)

	// set callback
	fwWin.SetCloseCallback(onStop)

	fwWin.SetFramebufferSizeCallback(resizeCallback)
	fwWin.SetCursorPosCallback(cursorPosCallback)
	fwWin.SetKeyCallback(keyEventCallback)
	fwWin.SetMouseButtonCallback(mouseButtonEventCallback)
	fwWin.SetCharCallback(runeCallback)

	monitor, videoMode = getMonitor(fwWin)

	de.SetWindowSize(de.win.Width, de.win.Height, de.win.IsFullScreen)

	//show
	fwWin.Show()

	de.sendLifecycle(lifecycle.StageFocused)
	canDoWork := worker.WorkAvailable()
	donec := make(chan struct{})
	go func() {
		mainLoop(de)
		close(donec)
	}()

	for {
		select {
		case <-donec:
			return
		case <-canDoWork:
			worker.DoWork()
		case <-de.drawDone:
			fwWin.SwapBuffers()
		}
		glfw.PollEvents()
	}

}

//SetWindowTitle ...
func (de *device) SetWindowTitle(title string) {
	fwWin.SetTitle(title)
}

//SetWindowSize ...
func (de *device) SetWindowSize(w, h int, fullScreen bool) {
	de.mu.Lock()
	de.mu.Unlock()

	if fullScreen {
		w = videoMode.Width
		h = videoMode.Height
		de.win.Width = w
		de.win.Height = h
		fwWin.SetMonitor(monitor, 0, 0, w, h, videoMode.RefreshRate)
		return
	}
	de.win.Width = w
	de.win.Height = h
	fwWin.SetSize(w, h)

}

// IsFullScreen ...
func (de *device) IsFullScreen() bool {
	de.mu.Lock()
	defer de.mu.Unlock()
	return de.win.IsFullScreen
}

func (de *device) SetFullScreen() {
	de.win.IsFullScreen = true
	de.SetWindowSize(0, 0, true)

}

func getMonitor(win *glfw.Window) (m *glfw.Monitor, v *glfw.VidMode) {
	mu.Lock()
	if m = win.GetMonitor(); m == nil {
		m = glfw.GetPrimaryMonitor()
	}

	v = m.GetVideoMode()

	return
}

func keyEventCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	k := convKeyCode(key)
	if action == glfw.Press {
		input.KeyPress(k)
	}
	if action == glfw.Release {
		input.KeyRelease(k)
	}

}

func mouseButtonEventCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	b := convMouse(button)
	if action == glfw.Press {
		input.MousePress(b)
	}
	if action == glfw.Release {
		input.MousePress(b)
	}
}

func runeCallback(w *glfw.Window, r rune) {
	input.SetRuns([]rune{r})
}

func cursorPosCallback(w *glfw.Window, x float64, y float64) {
	input.SetCursor(int(x), int(y))
}

func resizeCallback(win *glfw.Window, w, h int) {
	if w == ow && oh == h {
		return
	}
	ow, oh = w, h
	theDevice.win.Width = int(w)
	theDevice.win.Height = int(h)
	theDevice.sendEvents(reseizeEvent{})

}

//on window close ...
func onStop(win *glfw.Window) {
	theDevice.sendLifecycle(lifecycle.StageDead)
	theDevice.sendEvents(stopPumping{})
}

func convMouse(b glfw.MouseButton) input.MouseButton {

	switch b {
	case glfw.MouseButtonLeft:
		return input.MouseButtonLeft
	case glfw.MouseButtonRight:
		return input.MouseButtonRight
	case glfw.MouseButtonMiddle:
		return input.MouseButtonMiddle
	case glfw.MouseButton4:
		return input.MouseButton4
	case glfw.MouseButton5:
		return input.MouseButton5
	case glfw.MouseButton6:
		return input.MouseButton6
	case glfw.MouseButton7:
		return input.MouseButton7
	case glfw.MouseButton8:
		return input.MouseButton8
	}

	return input.MouseButtonNone

}

func convKeyCode(aKeyCode glfw.Key) input.Key {

	switch aKeyCode {

	case glfw.KeyEscape:
		return input.KeyEscape

	case glfw.KeyF1:
		return input.KeyF1
	case glfw.KeyF2:
		return input.KeyF2
	case glfw.KeyF3:
		return input.KeyF3
	case glfw.KeyF4:
		return input.KeyF4
	case glfw.KeyF5:
		return input.KeyF5
	case glfw.KeyF6:
		return input.KeyF6
	case glfw.KeyF7:
		return input.KeyF7
	case glfw.KeyF8:
		return input.KeyF8
	case glfw.KeyF9:
		return input.KeyF9
	case glfw.KeyF10:
		return input.KeyF10
	case glfw.KeyF11:
		return input.KeyF11
	case glfw.KeyF12:
		return input.KeyF12

	case glfw.Key1:
		return input.Key1
	case glfw.Key2:
		return input.Key2
	case glfw.Key3:
		return input.Key3
	case glfw.Key4:
		return input.Key4
	case glfw.Key5:
		return input.Key5
	case glfw.Key6:
		return input.Key6
	case glfw.Key7:
		return input.Key7
	case glfw.Key8:
		return input.Key8
	case glfw.Key9:
		return input.Key9
	case glfw.Key0:
		return input.Key0
	case glfw.KeyQ:
		return input.KeyQ
	case glfw.KeyW:
		return input.KeyW
	case glfw.KeyE:
		return input.KeyE
	case glfw.KeyR:
		return input.KeyR
	case glfw.KeyT:
		return input.KeyT
	case glfw.KeyY:
		return input.KeyY
	case glfw.KeyU:
		return input.KeyU
	case glfw.KeyI:
		return input.KeyI
	case glfw.KeyO:
		return input.KeyO
	case glfw.KeyP:
		return input.KeyP
	case glfw.KeyA:
		return input.KeyA
	case glfw.KeyS:
		return input.KeyS
	case glfw.KeyD:
		return input.KeyD
	case glfw.KeyF:
		return input.KeyF
	case glfw.KeyG:
		return input.KeyG
	case glfw.KeyH:
		return input.KeyH
	case glfw.KeyJ:
		return input.KeyJ
	case glfw.KeyK:
		return input.KeyK
	case glfw.KeyL:
		return input.KeyL
	case glfw.KeyZ:
		return input.KeyZ
	case glfw.KeyX:
		return input.KeyX
	case glfw.KeyC:
		return input.KeyC
	case glfw.KeyV:
		return input.KeyV
	case glfw.KeyB:
		return input.KeyB
	case glfw.KeyN:
		return input.KeyN
	case glfw.KeyM:
		return input.KeyM

	case glfw.KeyGraveAccent:
		return input.KeyGraveAccent
	case glfw.KeyMinus:
		return input.KeyMinus
	case glfw.KeyEqual:
		return input.KeyEqual
	case glfw.KeyBackspace:
		return input.KeyBackspace
	case glfw.KeyTab:
		return input.KeyTab
	case glfw.KeyLeftBracket:
		return input.KeyLeftSquareBracket
	case glfw.KeyRightBracket:
		return input.KeyRightSquareBracket
	case glfw.KeyBackslash:
		return input.KeyBackslash
	case glfw.KeyEnter:
		return input.KeyEnter
	case glfw.KeySemicolon:
		return input.KeySemicolon
	case glfw.KeyApostrophe:
		return input.KeyApostrophe
	case glfw.KeyComma:
		return input.KeyComma
	case glfw.KeyPeriod:
		return input.KeyFullStop
	case glfw.KeySlash:
		return input.KeySlash
	case glfw.KeySpace:
		return input.KeySpacebar

	// case 23:
	// return input.KeyCapsLock
	case glfw.KeyLeftShift:
		return input.KeyLeftShift
	case glfw.KeyRightShift:
		return input.KeyRightShift
	case glfw.KeyLeftControl:
		return input.KeyLeftCtrl
	case glfw.KeyLeftSuper:
		return input.KeyLeftGUI
	case glfw.KeyLeftAlt:
		return input.KeyLeftAlt
	case glfw.KeyRightAlt:
		return input.KeyRightAlt
	case glfw.KeyRightSuper:
		return input.KeyRightGUI
	case glfw.KeyMenu:
		return input.KeyMenu
	case glfw.KeyRightControl:
		return input.KeyRightCtrl

	case glfw.KeyPrintScreen:
		return input.KeyPrintScrn
	case glfw.KeyScrollLock:
		return input.KeyScrollLock
	// case 150:
	// return input.KeyInternational

	case glfw.KeyPause:
		return input.KeyPause
	case glfw.KeyInsert:
		return input.KeyInsert
	case glfw.KeyHome:
		return input.KeyHome
	case glfw.KeyPageUp:
		return input.KeyPageUp
	case glfw.KeyDelete:
		return input.KeyDelete
	case glfw.KeyEnd:
		return input.KeyEnd
	case glfw.KeyPageDown:
		return input.KeyPageDown

	case glfw.KeyUp:
		return input.KeyUpArrow
	case glfw.KeyLeft:
		return input.KeyLeftArrow
	case glfw.KeyDown:
		return input.KeyDownArrow
	case glfw.KeyRight:
		return input.KeyRightArrow

	case glfw.KeyNumLock:
		return input.KeyKpNumLock
	case glfw.KeyKPDivide:
		return input.KeyKpSlash
	case glfw.KeyKPMultiply:
		return input.KeyKpAsterisk
	case glfw.KeyKPSubtract:
		return input.KeyKpMinus
	case glfw.KeyKP7:
		return input.KeyKp7
	case glfw.KeyKP8:
		return input.KeyKp8
	case glfw.KeyKP9:
		return input.KeyKp9
	case glfw.KeyKPAdd:
		return input.KeyKpPlus
	case glfw.KeyKP4:
		return input.KeyKp4
	case glfw.KeyKP5:
		return input.KeyKp5
	case glfw.KeyKP6:
		return input.KeyKp6
	case glfw.KeyKP1:
		return input.KeyKp1
	case glfw.KeyKP2:
		return input.KeyKp2
	case glfw.KeyKP3:
		return input.KeyKp3
	case glfw.KeyKPEnter:
		return input.KeyKpEnter
	case glfw.KeyKP0:
		return input.KeyKp0
	case glfw.KeyKPDecimal:
		return input.KeyKpDot
	case glfw.KeyKPEqual:
		return input.KeyKpEqual

		// mobile
		// case 113:
		// return input.KeyMute
		// case 115:
		// return input.KeyVolumeUp
		// case 114:
		// return input.KeyVolumeDown

	}

	return input.KeyUnknown
}
