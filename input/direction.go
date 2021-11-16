package input

// Direction is the direction of the key(botton)
type Direction uint8

const (
	DirNone    Direction = 0
	DirPress   Direction = 1
	DirRelease Direction = 2

	// DirStep is a simultaneous press and release, such as a single step of a
	// mouse wheel.
	//
	// Its value equals DirPress | DirRelease.
	DirStep Direction = 3
)
