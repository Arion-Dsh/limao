//go:build linux || darwin
// +build linux darwin

package main

import (
	"fmt"
	"limao"
)

type app struct{}

func (app *app) PreLoad() {
	// limao.LoadDrawable(i)
}

func (app *app) Update() {

	if limao.IsKeyPressed(limao.Key1) {
		fmt.Println("111")
	}
}
func (app *app) Draw() {

}

func main() {
	a := new(app)

	o := &limao.RunOpts{Width: 600, Height: 800}
	limao.Run(a, o)

}
