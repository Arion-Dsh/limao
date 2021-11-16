package input

// MouseButton is a mouse button.
type MouseButton int32

// IsWheel reports whether the button is for a scroll wheel.
func (b MouseButton) IsWheel() bool {
	return b < 0
}

const (
	MouseButtonNone   MouseButton = +0
	MouseButtonLeft               = +1
	MouseButtonMiddle             = +2
	MouseButtonRight              = +3
	MouseButton4                  = +4
	MouseButton5                  = +5
	MouseButton6                  = +6
	MouseButton7                  = +7
	MouseButton8                  = +8

	MouseButtonWheelUp    = -1
	MouseButtonWheelDown  = -2
	MouseButtonWheelLeft  = -3
	MouseButtonWheelRight = -4
)

func MousePress(b MouseButton) {
	i.mu.Lock()
	i.mouse[b] = true
	i.mu.Unlock()
}

func MouseRelease(b MouseButton) {
	i.mu.Lock()
	delete(i.mouse, b)
	i.mu.Unlock()
}

func IsMousePressed(b MouseButton) bool {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.mouse == nil {
		return false
	}
	_, ok := i.mouse[b]
	return ok
}
