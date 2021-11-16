package limao

import "limao/input"

type Key uint32

func (k Key) String() string {
	s, ok := keyMap[k]
	if ok {
		return s
	}
	return "KeyUnknown"
}

var keyMap map[Key]string = map[Key]string{
	KeyUnknown: "KeyUnknown",

	KeyA: "KeyA",
	KeyB: "KeyB",
	KeyC: "KeyC",
	KeyD: "KeyD",
	KeyE: "KeyE",
	KeyF: "KeyF",
	KeyG: "KeyG",
	KeyH: "KeyH",
	KeyI: "KeyI",
	KeyJ: "KeyJ",
	KeyK: "KeyK",
	KeyL: "KeyL",
	KeyM: "KeyM",
	KeyN: "KeyN",
	KeyO: "KeyO",
	KeyP: "KeyP",
	KeyQ: "KeyQ",
	KeyR: "KeyR",
	KeyS: "KeyS",
	KeyT: "KeyT",
	KeyU: "KeyU",
	KeyV: "KeyV",
	KeyW: "KeyW",
	KeyX: "KeyX",
	KeyY: "KeyY",
	KeyZ: "KeyZ",

	Key1: "Key1",
	Key2: "Key2",
	Key3: "Key3",
	Key4: "Key4",
	Key5: "Key5",
	Key6: "Key6",
	Key7: "Key7",
	Key8: "Key8",
	Key9: "Key9",
	Key0: "Key0",

	KeyReturnEnter:        "KeyReturnEnter",
	KeyEscape:             "KeyEscape",
	KeyDeleteBackspace:    "KeyDeleteBackspace",
	KeyTab:                "KeyTab",
	KeySpacebar:           "KeySpacebar",
	KeyHyphenMinus:        "KeyHyphenMinus",        // -
	KeyEqualSign:          "KeyEqualSign",          // =
	KeyLeftSquareBracket:  "KeyLeftSquareBracket",  // [
	KeyRightSquareBracket: "KeyRightSquareBracket", // ]
	KeyBackslash:          "KeyBackslash",          // \
	KeySemicolon:          "KeySemicolon",          // ;
	KeyApostrophe:         "KeyApostrophe",         // '
	KeyGraveAccent:        "KeyGraveAccent",        // `
	KeyComma:              "KeyComma",              // ,
	KeyFullStop:           "KeyFullStop",           // .
	KeySlash:              "KeySlash",              // /
	KeyCapsLock:           "KeyCapsLock",

	KeyF1:  "KeyF1",
	KeyF2:  "KeyF2",
	KeyF3:  "KeyF3",
	KeyF4:  "KeyF4",
	KeyF5:  "KeyF5",
	KeyF6:  "KeyF6",
	KeyF7:  "KeyF7",
	KeyF8:  "KeyF8",
	KeyF9:  "KeyF9",
	KeyF10: "KeyF10",
	KeyF11: "KeyF11",
	KeyF12: "KeyF12",

	KeyPrintScrn:  "KeyPrintScrn",
	KeyScrollLock: "KeyScrollLock",

	KeyPause:         "KeyPause",
	KeyInsert:        "KeyInsert",
	KeyHome:          "KeyHome",
	KeyPageUp:        "KeyPageUp",
	KeyDeleteForward: "KeyDeleteForward",
	KeyEnd:           "KeyEnd",
	KeyPageDown:      "KeyPageDown",

	KeyRightArrow: "KeyRightArrow",
	KeyLeftArrow:  "KeyLeftArrow",
	KeyDownArrow:  "KeyDownArrow",
	KeyUpArrow:    "KeyUpArrow",

	KeyKPNumLock:     "KeyKPNumLock",
	KeyKPSlash:       "KeyKPSlash",       // /
	KeyKPAsterisk:    "KeyKPAsterisk",    // *
	KeyKPHyphenMinus: "KeyKPHyphenMinus", // -
	KeyKPPlusSign:    "KeyKPPlusSign",    // +
	KeyKPEnter:       "KeyKPEnter",
	KeyKP1:           "KeyKP1",
	KeyKP2:           "KeyKP2",
	KeyKP3:           "KeyKP3",
	KeyKP4:           "KeyKP4",
	KeyKP5:           "KeyKP5",
	KeyKP6:           "KeyKP6",
	KeyKP7:           "KeyKP7",
	KeyKP8:           "KeyKP8",
	KeyKP9:           "KeyKP9",
	KeyKP0:           "KeyKP0",
	KeyKPDot:         "KeyKPDot",   // .
	KeyKPEqual:       "KeyKPEqual", // =

	KeyHelp: "KeyHelp",

	KeyMute:       "KeyMute",
	KeyVolumeUp:   "KeyVolumeUp",
	KeyVolumeDown: "KeyVolumeDown",

	KeyLeftCtrl:   "KeyLeftCtrl",
	KeyLeftShift:  "KeyLeftShift",
	KeyLeftAlt:    "KeyLeftAlt",
	KeyLeftGUI:    "KeyLeftGUI",
	KeyRightCtrl:  "KeyRightCtrl",
	KeyRightShift: "KeyRightShift",
	KeyRightAlt:   "KeyRightAlt",
	KeyRightGUI:   "KeyRightGUI",

	KeyCompose: "KeyCompose",
}

const (
	KeyUnknown Key = Key(input.KeyUnknown)

	KeyA = Key(input.KeyA)
	KeyB = Key(input.KeyB)
	KeyC = Key(input.KeyC)
	KeyD = Key(input.KeyD)
	KeyE = Key(input.KeyE)
	KeyF = Key(input.KeyF)
	KeyG = Key(input.KeyG)
	KeyH = Key(input.KeyH)
	KeyI = Key(input.KeyI)
	KeyJ = Key(input.KeyJ)
	KeyK = Key(input.KeyK)
	KeyL = Key(input.KeyL)
	KeyM = Key(input.KeyM)
	KeyN = Key(input.KeyN)
	KeyO = Key(input.KeyO)
	KeyP = Key(input.KeyP)
	KeyQ = Key(input.KeyQ)
	KeyR = Key(input.KeyR)
	KeyS = Key(input.KeyS)
	KeyT = Key(input.KeyT)
	KeyU = Key(input.KeyU)
	KeyV = Key(input.KeyV)
	KeyW = Key(input.KeyW)
	KeyX = Key(input.KeyX)
	KeyY = Key(input.KeyY)
	KeyZ = Key(input.KeyZ)

	Key1 = Key(input.Key1)
	Key2 = Key(input.Key2)
	Key3 = Key(input.Key3)
	Key4 = Key(input.Key4)
	Key5 = Key(input.Key5)
	Key6 = Key(input.Key6)
	Key7 = Key(input.Key7)
	Key8 = Key(input.Key8)
	Key9 = Key(input.Key9)
	Key0 = Key(input.Key0)

	KeyReturnEnter        = Key(input.KeyEnter)
	KeyEscape             = Key(input.KeyEscape)
	KeyDeleteBackspace    = Key(input.KeyBackspace)
	KeyTab                = Key(input.KeyTab)
	KeySpacebar           = Key(input.KeySpacebar)
	KeyHyphenMinus        = Key(input.KeyMinus)              // -
	KeyEqualSign          = Key(input.KeyEqual)              // =
	KeyLeftSquareBracket  = Key(input.KeyLeftSquareBracket)  // [
	KeyRightSquareBracket = Key(input.KeyRightSquareBracket) // ]
	KeyBackslash          = Key(input.KeyBackslash)          // \
	KeySemicolon          = Key(input.KeySemicolon)          // ;
	KeyApostrophe         = Key(input.KeyApostrophe)         // '
	KeyGraveAccent        = Key(input.KeyGraveAccent)        // `
	KeyComma              = Key(input.KeyComma)              // ,
	KeyFullStop           = Key(input.KeyFullStop)           // .
	KeySlash              = Key(input.KeySlash)              // /
	KeyCapsLock           = Key(input.KeyCapsLock)

	KeyF1  = Key(input.KeyF1)
	KeyF2  = Key(input.KeyF2)
	KeyF3  = Key(input.KeyF3)
	KeyF4  = Key(input.KeyF4)
	KeyF5  = Key(input.KeyF5)
	KeyF6  = Key(input.KeyF6)
	KeyF7  = Key(input.KeyF7)
	KeyF8  = Key(input.KeyF8)
	KeyF9  = Key(input.KeyF9)
	KeyF10 = Key(input.KeyF10)
	KeyF11 = Key(input.KeyF11)
	KeyF12 = Key(input.KeyF12)

	KeyPrintScrn  = Key(input.KeyPrintScrn)
	KeyScrollLock = Key(input.KeyScrollLock)

	KeyPause         = Key(input.KeyPause)
	KeyInsert        = Key(input.KeyInsert)
	KeyHome          = Key(input.KeyHome)
	KeyPageUp        = Key(input.KeyPageUp)
	KeyDeleteForward = Key(input.KeyDelete)
	KeyEnd           = Key(input.KeyEnd)
	KeyPageDown      = Key(input.KeyPageDown)

	KeyRightArrow = Key(input.KeyRightArrow)
	KeyLeftArrow  = Key(input.KeyLeftArrow)
	KeyDownArrow  = Key(input.KeyDownArrow)
	KeyUpArrow    = Key(input.KeyUpArrow)

	KeyKPNumLock     = Key(input.KeyKpNumLock)
	KeyKPSlash       = Key(input.KeyKpSlash)    // /
	KeyKPAsterisk    = Key(input.KeyKpAsterisk) // *
	KeyKPHyphenMinus = Key(input.KeyKpMinus)    // -
	KeyKPPlusSign    = Key(input.KeyKpPlus)     // +
	KeyKPEnter       = Key(input.KeyKpEnter)
	KeyKP1           = Key(input.KeyKp1)
	KeyKP2           = Key(input.KeyKp2)
	KeyKP3           = Key(input.KeyKp3)
	KeyKP4           = Key(input.KeyKp4)
	KeyKP5           = Key(input.KeyKp5)
	KeyKP6           = Key(input.KeyKp6)
	KeyKP7           = Key(input.KeyKp7)
	KeyKP8           = Key(input.KeyKp8)
	KeyKP9           = Key(input.KeyKp9)
	KeyKP0           = Key(input.KeyKp0)
	KeyKPDot         = Key(input.KeyKpDot)   // .
	KeyKPEqual       = Key(input.KeyKpEqual) // =

	KeyHelp = Key(input.KeyHelp)

	KeyMute       = Key(input.KeyMute)
	KeyVolumeUp   = Key(input.KeyVolumeUp)
	KeyVolumeDown = Key(input.KeyVolumeDown)

	KeyLeftCtrl   = Key(input.KeyLeftCtrl)
	KeyLeftShift  = Key(input.KeyLeftShift)
	KeyLeftAlt    = Key(input.KeyLeftAlt)
	KeyLeftGUI    = Key(input.KeyLeftGUI)
	KeyRightCtrl  = Key(input.KeyRightCtrl)
	KeyRightShift = Key(input.KeyRightShift)
	KeyRightAlt   = Key(input.KeyRightAlt)
	KeyRightGUI   = Key(input.KeyRightGUI)

	KeyCompose = Key(input.KeyCompose)
)
