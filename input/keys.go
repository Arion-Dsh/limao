package input

const (
	KeyUnknown            Key = 0
	KeyEscape                 = 4   //9  Esc
	KeyF1                     = 5   // 67 F1
	KeyF2                     = 6   //68 F2
	KeyF3                     = 7   // 69   F3
	KeyF4                     = 8   // 70   F4
	KeyF5                     = 9   // 71   F5
	KeyF6                     = 10  // 72   F6
	KeyF7                     = 11  //73   F7
	KeyF8                     = 12  // 74   F8
	KeyF9                     = 13  //75   F9
	KeyF10                    = 14  //76   F10
	KeyF11                    = 15  //95  F11
	KeyF12                    = 16  //96  F12
	KeyPrintScrn              = 17  //111 PrintScrn
	KeyScrollLock             = 18  //78  Scroll Lock
	KeyPause                  = 19  // 110 Pause
	KeyGraveAccent            = 20  //49  `
	Key1                      = 21  //10  1
	Key2                      = 22  //11  2
	Key3                      = 23  // 12 3
	Key4                      = 24  //13  4
	Key5                      = 25  //14  5
	Key6                      = 26  // 15 6
	Key7                      = 27  //16  7
	Key8                      = 28  // 17 8
	Key9                      = 29  //18 9
	Key0                      = 30  //19 0
	KeyMinus                  = 31  //31 -
	KeyEqual                  = 32  //21 =
	KeyBackspace              = 33  // 22 Backspace
	KeyInsert                 = 34  //106 Insert
	KeyHome                   = 35  //97  Home
	KeyPageUp                 = 36  //99  Page Up
	KeyKpNumLock              = 37  // 77 Num Lock
	KeyKpSlash                = 39  //112 Kp /
	KeyKpAsterisk             = 40  //63  Kp *
	KeyKpMinus                = 41  //82  Kp -
	KeyTab                    = 42  //23  Tab
	KeyQ                      = 43  //24  Q
	KeyW                      = 44  //25  W
	KeyE                      = 45  //26  // E
	KeyR                      = 46  //27  // R
	KeyT                      = 47  //28  // T
	KeyY                      = 48  //29  // Y
	KeyU                      = 49  //30  // U
	KeyI                      = 50  //31  // I
	KeyO                      = 51  // 32  // O
	KeyP                      = 52  //33  // P
	KeyLeftSquareBracket      = 53  //34  // [
	KeyRightSquareBracket     = 54  //35  // ]
	KeyEnter                  = 55  //36  // Enter
	KeyDelete                 = 56  //107 // Delete
	KeyEnd                    = 57  //103 // End
	KeyPageDown               = 58  //105 // Page Down
	KeyKp7                    = 59  // Kp 7
	KeyKp8                    = 60  //80  // Kp 8
	KeyKp9                    = 61  //81  // Kp 9
	KeyKpPlus                 = 62  //86  // Kp +
	KeyCapsLock               = 63  //66  // Caps Lock
	KeyA                      = 64  //38  // A
	KeyS                      = 65  //39  // S
	KeyD                      = 66  // 40  // D
	KeyF                      = 67  //41  // F
	KeyG                      = 68  //42  // G
	KeyH                      = 69  //43  // H
	KeyJ                      = 70  //44  // J
	KeyK                      = 71  //45  // K
	KeyL                      = 72  //46  // L
	KeySemicolon              = 73  //47  // ;
	KeyApostrophe             = 74  //48  // '
	KeyKp4                    = 75  //83  // Kp 4
	KeyKp5                    = 76  //84  // Kp 5
	KeyKp6                    = 77  //85  // Kp 6
	KeyLeftShift              = 78  //50  // Shift Left
	KeyInternational          = 79  //94  // International
	KeyZ                      = 80  // 52  // Z
	KeyX                      = 81  //53  // X
	KeyC                      = 82  //54  // C
	KeyV                      = 83  // 55  // V
	KeyB                      = 84  // 56  // B
	KeyN                      = 85  //57  // N
	KeyM                      = 86  //58  // M
	KeyComma                  = 87  //59  // ,
	KeyFullStop               = 88  //60  // .
	KeySlash                  = 89  //61  // /
	KeyRightShift             = 90  //62  // Shift Right
	KeyBackslash              = 91  //51  // \
	KeyUpArrow                = 92  //98  // Cursor Up
	KeyKp1                    = 93  //87  // Kp 1
	KeyKp2                    = 94  // 88  // Kp 2
	KeyKp3                    = 95  //89  // Kp 3
	KeyKpEnter                = 96  //108 // Kp Enter
	KeyLeftCtrl               = 97  //37  // Ctrl Left
	KeyLeftGUI                = 98  //115 // Logo Left (-> Option)
	KeyLeftAlt                = 99  //64  // Alt Left (-> Command)
	KeySpacebar               = 100 //65  // Space
	KeyRightAlt               = 101 //113 // Alt Right (-> Command)
	KeyRightGUI               = 102 //116 // Logo Right (-> Option)
	KeyMenu                   = 103 //117 // Menu (-> International)
	KeyRightCtrl              = 104 // 109 // Ctrl Right
	KeyLeftArrow              = 105 //100 // Cursor Left
	KeyDownArrow              = 106 //104 // Cursor Down
	KeyRightArrow             = 107 // 102 // Cursor Right
	KeyKp0                    = 108 //90  // Kp 0
	KeyKpDot                  = 109 //91  // Kp .

	// mobile key
	KeyKpEqual = 110 // =
	KeyHelp    = 111

	KeyMute       = 112
	KeyVolumeUp   = 113
	KeyVolumeDown = 114

	// The following codes are not part of the standard USB HID Usage IDs for
	// keyboards. See http://www.usb.org/developers/hidpage/Hut1_12v2.pdf
	//
	// Usage IDs are uint16s, so these non-standard values start at 0x10000.

	// KeyCompose is the Key for a compose key, sometimes called a multi key,
	// used to input non-ASCII characters such as ñ being composed of n and ~.
	//
	// See https://en.wikipedia.org/wiki/Compose_key
	KeyCompose = 0x10000
)
