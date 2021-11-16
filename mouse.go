package limao

import "limao/input"

type MouseButton input.MouseButton

const (
	MouseButtonNone   MouseButton = MouseButton(input.MouseButtonNone)
	MouseButtonLeft               = MouseButton(input.MouseButtonLeft)
	MouseButtonMiddle             = MouseButton(input.MouseButtonMiddle)
	MouseButtonRight              = MouseButton(input.MouseButtonRight)
	MouseButton4                  = MouseButton(input.MouseButton4)
	MouseButton5                  = MouseButton(input.MouseButton5)
	MouseButton6                  = MouseButton(input.MouseButton6)
	MouseButton7                  = MouseButton(input.MouseButton7)
	MouseButton8                  = MouseButton(input.MouseButton8)

	MouseButtonWheelUp    = MouseButton(input.MouseButtonWheelUp)
	MouseButtonWheelDown  = MouseButton(input.MouseButtonWheelDown)
	MouseButtonWheelLeft  = MouseButton(input.MouseButtonWheelLeft)
	MouseButtonWheelRight = MouseButton(input.MouseButtonWheelRight)
)
