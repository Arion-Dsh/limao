//go:build js
// +build js

package device

import (
	"limao/device/lifecycle"
	"limao/graphics"
	"limao/graphics/gl"
	"limao/input"
	"syscall/js"
)

var (
	jsDoc    = js.Global().Get("document")
	canvasEl js.Value
)

func (de *device) run() {
	if !jsDoc.Truthy() {
		return
	}

	bodyStyle := jsDoc.Get("body").Get("style")
	bodyStyle.Set("margin", "0")

	canvasEl = jsDoc.Call("createElement", "canvas")
	jsDoc.Get("body").Call("appendChild", canvasEl)

	ctx, _ := gl.NewContext(canvasEl)
	de.ctx = ctx

	// add event listener
	de.processEvents(canvasEl)

	//draw loop
	donec := make(chan struct{})

	go func() {
		mainLoop(de)
		close(donec)
	}()

	de.sendLifecycle(lifecycle.StageFocused)
	js.Global().Call("requestAnimationFrame", de.render())
	<-donec

}

var vw, vh float64

func (de *device) render() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		<-de.drawDone

		// set canvas size
		width := jsDoc.Get("body").Get("clientWidth").Float()
		height := jsDoc.Get("body").Get("clientHeight").Float()

		if vw != width && vh != height {
			canvasEl.Set("width", width)
			canvasEl.Set("height", height)
			graphics.SetViewPort(int(width), int(height))
		}

		js.Global().Call("requestAnimationFrame", de.render())
		return nil
	})

}

func (de *device) processEvents(doc js.Value) {

	// mouse move

	doc.Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		x := e.Get("clientX").Int()
		y := e.Get("clientY").Int()
		input.SetCursor(int(x), int(y))
		return nil
	}))

	doc.Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		k := e.Get("key").String()
		l := e.Get("location").Int()
		input.KeyPress(convKeyCode(k, l))
		return nil
	}))

	doc.Call("addEventListener", "keyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		k := e.Get("key").String()
		l := e.Get("location").Int()
		input.KeyRelease(convKeyCode(k, l))
		return nil
	}))

	doc.Call("addEventListener", "mousedown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		b := e.Get("button").Int()
		input.MousePress(convMouse(b))
		return nil
	}))

	doc.Call("addEventListener", "mouseup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		b := e.Get("button").Int()
		input.MousePress(convMouse(b))
		return nil
	}))

}

//SetWindowSize ...
func (de *device) SetWindowSize(w, h int32) {
	return
}

func convMouse(b int) input.MouseButton {

	switch b {
	case 0:
		return input.MouseButtonLeft
	case 2:
		return input.MouseButtonRight
	case 1:
		return input.MouseButtonMiddle
	}

	return input.MouseButtonNone

}

func convKeyCode(aKey string, location int) input.Key {

	switch aKey {

	case "Escape", "Esc":
		return input.KeyEscape

	case "F1":
		return input.KeyF1
	case "F2":
		return input.KeyF2
	case "F3":
		return input.KeyF3
	case "F4":
		return input.KeyF4
	case "F5":
		return input.KeyF5
	case "F6":
		return input.KeyF6
	case "F7":
		return input.KeyF7
	case "F8":
		return input.KeyF8
	case "F9":
		return input.KeyF9
	case "F10":
		return input.KeyF10
	case "F11":
		return input.KeyF11
	case "F12":
		return input.KeyF12

	case "1", "!":
		return input.Key1
	case "2", "@":
		return input.Key2
	case "3", "#":
		return input.Key3
	case "4", "$":
		return input.Key4
	case "5", "%":
		return input.Key5
	case "6", "^":
		return input.Key6
	case "7", "&":
		return input.Key7
	case "8", "*":
		return input.Key8
	case "9", "(":
		return input.Key9
	case "0", ")":
		return input.Key0
	case "q", "Q":
		return input.KeyQ
	case "w", "W":
		return input.KeyW
	case "e", "E":
		return input.KeyE
	case "r", "R":
		return input.KeyR
	case "t", "T":
		return input.KeyT
	case "y", "Y":
		return input.KeyY
	case "u", "U":
		return input.KeyU
	case "i", "I":
		return input.KeyI
	case "o", "O":
		return input.KeyO
	case "p", "P":
		return input.KeyP
	case "a", "A":
		return input.KeyA
	case "s", "S":
		return input.KeyS
	case "d", "D":
		return input.KeyD
	case "f", "F":
		return input.KeyF
	case "g", "G":
		return input.KeyG
	case "h", "H":
		return input.KeyH
	case "j", "J":
		return input.KeyJ
	case "k", "K":
		return input.KeyK
	case "l", "L":
		return input.KeyL
	case "z", "Z":
		return input.KeyZ
	case "x", "X":
		return input.KeyX
	case "c", "C":
		return input.KeyC
	case "v", "V":
		return input.KeyV
	case "b", "B":
		return input.KeyB
	case "n", "N":
		return input.KeyN
	case "m", "M":
		return input.KeyM

	case "`", "~":
		return input.KeyGraveAccent
	case "-", "_":
		return input.KeyMinus
	case "=", "+":
		return input.KeyEqual
	case "Backspace":
		return input.KeyBackspace
	case "Tab":
		return input.KeyTab
	case "[", "{":
		return input.KeyLeftSquareBracket
	case "]", "}":
		return input.KeyRightSquareBracket
	case "\\", "|":
		return input.KeyBackslash
	case "Enter":
		return input.KeyEnter
	case ";", ":":
		return input.KeySemicolon
	case "'", "\"":
		return input.KeyApostrophe
	case ",", "<":
		return input.KeyComma
	case ".", ">":
		return input.KeyFullStop
	case "/", "?":
		return input.KeySlash
	case " ", "Spacebar":
		return input.KeySpacebar

	case "CapsLock":
		return input.KeyCapsLock
	case "Shift":
		if location == 1 {
			return input.KeyLeftShift
		}
		return input.KeyRightShift
	case "Control":
		if location == 1 {
			return input.KeyLeftCtrl
		}
		return input.KeyRightCtrl
	case "Meta", "OS":
		if location == 1 {
			return input.KeyLeftGUI
		}
		return input.KeyRightGUI
	case "Alt":
		if location == 1 {
			return input.KeyLeftAlt
		}
		return input.KeyRightAlt
	case "ContextMenu":
		return input.KeyMenu

	case "PrintScreen":
		return input.KeyPrintScrn
	case "ScrollLock", "Scroll":
		return input.KeyScrollLock
	// case 150:
	// return input.KeyInternational

	case "Pause":
		return input.KeyPause
	case "Insert":
		return input.KeyInsert
	case "Home":
		return input.KeyHome
	case "PageUp":
		return input.KeyPageUp
	case "Delete":
		return input.KeyDelete
	case "End":
		return input.KeyEnd
	case "PageDown":
		return input.KeyPageDown

	case "ArrowUp", "Up":
		return input.KeyUpArrow
	case "ArrowLeft", "Left":
		return input.KeyLeftArrow
	case "ArrowDown", "Down":
		return input.KeyDownArrow
	case "ArrowRight", "Right":
		return input.KeyRightArrow

	// case "NumLock":
	// return input.KeyKpNumLock
	// case key.CodeKeypadSlash:
	// return input.KeyKpSlash
	// case key.CodeKeypadAsterisk:
	// return input.KeyKpAsterisk
	// case key.CodeKeypadHyphenMinus:
	// return input.KeyKpMinus
	// case key.CodeKeypad7:
	// return input.KeyKp7
	// case key.CodeKeypad8:
	// return input.KeyKp8
	// case key.CodeKeypad9:
	// return input.KeyKp9
	// case key.CodeKeypadPlusSign:
	// return input.KeyKpPlus
	// case key.CodeKeypad4:
	// return input.KeyKp4
	// case key.CodeKeypad5:
	// return input.KeyKp5
	// case key.CodeKeypad6:
	// return input.KeyKp6
	// case key.CodeKeypad1:
	// return input.KeyKp1
	// case key.CodeKeypad2:
	// return input.KeyKp2
	// case key.CodeKeypad3:
	// return input.KeyKp3
	// case key.CodeKeypadEnter:
	// return input.KeyKpEnter
	// case key.CodeKeypad0:
	// return input.KeyKp0
	// case key.CodeKeypadFullStop:
	// return input.KeyKpDot

	// // mobile
	// case key.CodeKeypadEqualSign:
	// return input.KeyKpEqual

	case "AudioVolumeMute", "VolumeMute":
		return input.KeyMute
	case "VolumeUp", "AudioVolumeUp":
		return input.KeyVolumeUp
	case "VolumeDown", "AudioVolumeDown":
		return input.KeyVolumeDown
	case "Compose":
		return input.KeyCompose

	}

	return input.KeyUnknown
}
