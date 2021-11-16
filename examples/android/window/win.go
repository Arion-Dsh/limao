//go:build linux
// +build linux

package main

import (
	"limao"
)

type app struct{}

func (app *app) PreLoad() {
}

func (app *app) Update() {

}
func (app *app) Draw() {
}

func main() {
	a := new(app)

	o := &limao.RunOpts{Width: 600, Height: 800}
	limao.Run(a, o)

}
