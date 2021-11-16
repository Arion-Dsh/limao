//go:build (linux || darwin || android) && !windows
// +build linux darwin android
// +build !windows

package gl

import gles "golang.org/x/mobile/gl"

func getVersion() float32 {

	v := float32(3.0)
	switch gles.Version() {
	case "GL_ES_2_0":
		v = 2.0
	case "GL_ES_3_0":
		v = 3.0
	case "GL_ES_3_1":
		v = 3.1
	}
	return v

}
