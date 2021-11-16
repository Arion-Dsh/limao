package m32

import "math"

// ScreenToGLCoords transforms from pixel coordinates to GL coordinates.
//
// This assumes that your pixel coordinate system considers its origin to be in the top left corner (GL's is in the bottom left).
// The coordinates x and y may be out of the range [0,screenWidth-1] and [0,screeneHeight-1].
//
// GL's coordinate system maps [screenWidth-1,0] to [1.0,1.0] and [0,screenHeight-1] to [-1.0,-1.0]. If x and y are out of the range, they'll still
// be mapped correctly, just off the screen. (e.g. if y = 2*(screenHeight-1) you'll get -3.0 for yOut)
//
// This is similar to Unproject, except for 2D cases and much simpler (especially since an inverse may always be found)
func ScreenToGLCoords(x, y int, screenWidth, screenHeight int) (xOut, yOut float32) {
	xOut = 2.0*float32(x)/float32(screenWidth-1) - 1.0
	yOut = -2.0*float32(y)/float32(screenHeight-1) + 1.0

	return
}

// GLToScreenCoords transforms from GL's proportional system to pixel coordinates.
//
// Assumes the pixel coordinate system has its origin in the top left corner. (GL's is in the bottom left)
//
// GL's coordinate system maps [screenWidth-1,0] to [1.0,1.0] and [0,screenHeight-1] to [-1.0,-1.0]. If x and y are out of the range, they'll still
// be mapped correctly, just off the screen. (e.g. if y=-3.0, you'll get 2*(screenHeight-1) for yOut)
//
// This is similar to Project, except for 2D cases and much simpler
func GLToScreenCoords(x, y float32, screenWidth, screenHeight int) (xOut, yOut int) {
	xOut = int((x + 1.0) * float32(screenWidth-1) / 2.0)
	yOut = int((1.0 - y) * float32(screenHeight-1) / 2.0)

	return
}

// Circle generates a circle centered at (0,0) with a given radius.
// The radii are assumed to be in GL's coordinate sizing.
//
// Technically this draws an ellipse with two axes that match with the X and Y axes, the reason it has a radiusX and radiusY is because GL's coordinate system
// is proportional to screen width and screen height. So if you have a non-square viewport, a single radius will appear
// to "squash" the circle in one direction (usually the Y direction), so the X and Y radius allow for a circle to be made.
// A good way to get the correct radii is with mathgl.ScreenToGLCoords(radius, radius, screenWidth, screenHeight) which will get you the correct
// proportional GL coords.
//
// The numSlices argument specifies how many triangles you want your circle divided into, setting this
// number to too low a value may cause problem (and too high will cause it to take a lot of memory and time to compute
// without much gain in resolution).
//
// This uses discrete triangles, not a triangle fan
func Circle(radiusX, radiusY float32, numSlices int) []Vec2 {
	twoPi := float32(2.0 * math.Pi)

	circlePoints := make([]Vec2, 0, numSlices*3)
	center := Vec2{0.0, 0.0}
	previous := Vec2{radiusX, 0.0}

	for theta := twoPi / float32(numSlices); !FloatEqual(theta, twoPi); theta = Clamp(theta+twoPi/float32(numSlices), 0.0, twoPi) {
		sin, cos := math.Sincos(float64(theta))
		curr := Vec2{float32(cos) * radiusX, float32(sin) * radiusY}

		circlePoints = append(circlePoints, center, previous, curr)
		previous = curr
	}

	// Now add the final point at theta=2pi
	circlePoints = append(circlePoints, center, previous, Vec2{radiusX, 0.0})
	return circlePoints
}

// Rect generates a 2-triangle rectangle for use with GL_TRIANGLES. The width and height should use GL's proportions (that is, where a width of 1.0
// is equivalent to half of the width of the render target); however, the y-coordinates grow downwards, not upwards. That is, it
// assumes you want the origin of the rectangle with the top-left corner at (0.0,0.0).
//
// Keep in mind that GL's coordinate system is proportional, so width=height will not result in a square unless your viewport is square.
// If you want to maintain proportionality regardless of screen size, use the results of w,h := ScreenToGLCoordsf(absoluteWidth, absoluteHeight, screenWidth, screenHeight);
// w,h=w+1,h-1 in the call to this function. (The w+1,h-1 step maps the coordinates to start at 0.0 rather than -1.0)
func Rect(width, height float32) []Vec2 {
	return []Vec2{
		{0.0, 0.0},
		{0.0, -height},
		{width, -height},

		{0.0, 0.0},
		{width, -height},
		{width, 0.0},
	}
}
