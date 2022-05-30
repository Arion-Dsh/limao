//go:build android || ios || mobile
// +build android ios mobile

package device

import (
	"limao/graphics"
	"limao/graphics/gl"
	"limao/input"
	"unicode"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/key"
	lifec "golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
)

func (de *device) run() {

	app.Main(func(a app.App) {

		var ctx gl.Context

		var vw, vh int = 4, 4

		go de.update()
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifec.Event:
				switch e.Crosses(lifec.StageVisible) {
				case lifec.CrossOn:
					ctx = gl.NewContextWithES(e.DrawContext)
					de.ctx = ctx
					graphics.Load(de.ctx)
					de.app.OnStart()
					de.start.Broadcast()
					a.Send(paint.Event{})
				case lifec.CrossOff:
					graphics.Release()
					de.app.OnStop()
					de.ctx = nil
				}
			case size.Event:
				if vw != e.WidthPx && vh != e.HeightPx {
					vw = e.WidthPx
					vh = e.HeightPx
					de.win.Width = vw
					de.win.Height = vh
					de.resize = true
				}
			case paint.Event:
				if de.ctx == nil || e.External {
					continue
				}
				if de.resize {
					graphics.SetViewPort(vw, vh)
					a.Send(paint.Event{})
					de.resize = false
					continue
				}

				graphics.Clear()
				de.app.Draw()
				graphics.Draw()
				a.Publish()
				a.Send(paint.Event{}) // keep animating
			case touch.Event:
				id := int(e.Sequence)
				if e.Type == touch.TypeEnd {
					input.TouchEnd(id)
				} else {
					input.TouchMove(id, int(e.X), int(e.Y))
				}
			case key.Event:
				k := convKeyCode(e.Code)
				switch e.Direction {
				case key.DirPress, key.DirNone:
					input.KeyPress(k)
				case key.DirRelease:
					input.KeyRelease(k)
				}
				// rune
				if e.Rune != -1 && unicode.IsPrint(e.Rune) {
					input.SetRuns([]rune{e.Rune})
				}

			}
		}
	})
}

func convKeyCode(aKeyCode key.Code) input.Key {

	switch aKeyCode {

	case key.CodeEscape:
		return input.KeyEscape

	case key.CodeF1:
		return input.KeyF1
	case key.CodeF2:
		return input.KeyF2
	case key.CodeF3:
		return input.KeyF3
	case key.CodeF4:
		return input.KeyF4
	case key.CodeF5:
		return input.KeyF5
	case key.CodeF6:
		return input.KeyF6
	case key.CodeF7:
		return input.KeyF7
	case key.CodeF8:
		return input.KeyF8
	case key.CodeF9:
		return input.KeyF9
	case key.CodeF10:
		return input.KeyF10
	case key.CodeF11:
		return input.KeyF11
	case key.CodeF12:
		return input.KeyF12

	case key.Code1:
		return input.Key1
	case key.Code2:
		return input.Key2
	case key.Code3:
		return input.Key3
	case key.Code4:
		return input.Key4
	case key.Code5:
		return input.Key5
	case key.Code6:
		return input.Key6
	case key.Code7:
		return input.Key7
	case key.Code8:
		return input.Key8
	case key.Code9:
		return input.Key9
	case key.Code0:
		return input.Key0
	case key.CodeQ:
		return input.KeyQ
	case key.CodeW:
		return input.KeyW
	case key.CodeE:
		return input.KeyE
	case key.CodeR:
		return input.KeyR
	case key.CodeT:
		return input.KeyT
	case key.CodeY:
		return input.KeyY
	case key.CodeU:
		return input.KeyU
	case key.CodeI:
		return input.KeyI
	case key.CodeO:
		return input.KeyO
	case key.CodeP:
		return input.KeyP
	case key.CodeA:
		return input.KeyA
	case key.CodeS:
		return input.KeyS
	case key.CodeD:
		return input.KeyD
	case key.CodeF:
		return input.KeyF
	case key.CodeG:
		return input.KeyG
	case key.CodeH:
		return input.KeyH
	case key.CodeJ:
		return input.KeyJ
	case key.CodeK:
		return input.KeyK
	case key.CodeL:
		return input.KeyL
	case key.CodeZ:
		return input.KeyZ
	case key.CodeX:
		return input.KeyX
	case key.CodeC:
		return input.KeyC
	case key.CodeV:
		return input.KeyV
	case key.CodeB:
		return input.KeyB
	case key.CodeN:
		return input.KeyN
	case key.CodeM:
		return input.KeyM

	case key.CodeGraveAccent:
		return input.KeyGraveAccent
	case key.CodeHyphenMinus:
		return input.KeyMinus
	case key.CodeEqualSign:
		return input.KeyEqual
	case key.CodeDeleteBackspace:
		return input.KeyBackspace
	case key.CodeTab:
		return input.KeyTab
	case key.CodeLeftSquareBracket:
		return input.KeyLeftSquareBracket
	case key.CodeRightSquareBracket:
		return input.KeyRightSquareBracket
	case key.CodeBackslash:
		return input.KeyBackslash
	case key.CodeReturnEnter:
		return input.KeyEnter
	case key.CodeSemicolon:
		return input.KeySemicolon
	case key.CodeApostrophe:
		return input.KeyApostrophe
	case key.CodeComma:
		return input.KeyComma
	case key.CodeFullStop:
		return input.KeyFullStop
	case key.CodeSlash:
		return input.KeySlash
	case key.CodeSpacebar:
		return input.KeySpacebar

	case key.CodeCapsLock:
		return input.KeyCapsLock
	case key.CodeLeftShift:
		return input.KeyLeftShift
	case key.CodeRightShift:
		return input.KeyRightShift
	case key.CodeLeftControl:
		return input.KeyLeftCtrl
	case key.CodeLeftGUI:
		return input.KeyLeftGUI
	case key.CodeLeftAlt:
		return input.KeyLeftAlt
	case key.CodeRightAlt:
		return input.KeyRightAlt
	case key.CodeRightGUI:
		return input.KeyRightGUI
	// case 126:
	// return input.KeyMenu
	case key.CodeRightControl:
		return input.KeyRightCtrl

	// case 99:
	// return input.KeyPrintScrn
	// case 70:
	// return input.KeyScrollLock
	// case 150:
	// return input.KeyInternational

	case key.CodePause:
		return input.KeyPause
	case key.CodeInsert:
		return input.KeyInsert
	case key.CodeHome:
		return input.KeyHome
	case key.CodePageUp:
		return input.KeyPageUp
	case key.CodeDeleteForward:
		return input.KeyDelete
	case key.CodeEnd:
		return input.KeyEnd
	case key.CodePageDown:
		return input.KeyPageDown

	case key.CodeUpArrow:
		return input.KeyUpArrow
	case key.CodeLeftArrow:
		return input.KeyLeftArrow
	case key.CodeDownArrow:
		return input.KeyDownArrow
	case key.CodeRightArrow:
		return input.KeyRightArrow

	case key.CodeKeypadNumLock:
		return input.KeyKpNumLock
	case key.CodeKeypadSlash:
		return input.KeyKpSlash
	case key.CodeKeypadAsterisk:
		return input.KeyKpAsterisk
	case key.CodeKeypadHyphenMinus:
		return input.KeyKpMinus
	case key.CodeKeypad7:
		return input.KeyKp7
	case key.CodeKeypad8:
		return input.KeyKp8
	case key.CodeKeypad9:
		return input.KeyKp9
	case key.CodeKeypadPlusSign:
		return input.KeyKpPlus
	case key.CodeKeypad4:
		return input.KeyKp4
	case key.CodeKeypad5:
		return input.KeyKp5
	case key.CodeKeypad6:
		return input.KeyKp6
	case key.CodeKeypad1:
		return input.KeyKp1
	case key.CodeKeypad2:
		return input.KeyKp2
	case key.CodeKeypad3:
		return input.KeyKp3
	case key.CodeKeypadEnter:
		return input.KeyKpEnter
	case key.CodeKeypad0:
		return input.KeyKp0
	case key.CodeKeypadFullStop:
		return input.KeyKpDot

	// mobile
	case key.CodeKeypadEqualSign:
		return input.KeyKpEqual
	case key.CodeMute:
		return input.KeyMute
	case key.CodeVolumeUp:
		return input.KeyVolumeUp
	case key.CodeVolumeDown:
		return input.KeyVolumeDown
	case key.CodeCompose:
		return input.KeyCompose

	}

	return input.KeyUnknown
}
