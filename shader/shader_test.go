package shader

import (
	"fmt"
	"testing"
)

func TestShader(t *testing.T) {

	opt := &ShaderOption{}
	opt.Uniforms = map[string]string{
		"p": "vec2",
		"v": "vec3",
	}
	opt.Fns = []string{`
vec2 a() {
	return vec2(0, 0);
}
	`}

	opt.FragColor = "gl_FragColor = gl_FragColor;"

	fs, _ := formatfs(opt)

	fmt.Print(fs)
}
