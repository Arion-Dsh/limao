//go:build js
// +build js

package main

import (
	"limao"
)

type app struct{}

func (app *app) OnStart() {}
func (app *app) OnStop()  {}

func (app *app) Update() {

}
func (app *app) Draw() {

}

func main() {

	a := new(app)
	o := &limao.RunOpts{Width: 1600, Height: 1800}
	limao.Run(a, o)

}
