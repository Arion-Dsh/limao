package limao

import "limao/input"

func IsKeyPressed(k Key) bool {
	return input.IsKeyPressed(input.Key(k))
}

func IsMousePressed(b MouseButton) bool {
	return input.IsMousePressed(input.MouseButton(b))
}
