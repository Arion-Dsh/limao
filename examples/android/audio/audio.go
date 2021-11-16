package main

import (
	"limao"
	"limao/audio"
)

var f limao.Asset
var f2 limao.Asset

var s *audio.Source
var s1 *audio.Source

type app struct {
	played bool
}

func (app *app) PreLoad() {
	// f, _ = os.Open("/home/arion/Downloads/Lemon.wav")
	f, _ = limao.OpenAsset("12_1.wav")
	f2, _ = limao.OpenAsset("123_1.wav")
	audio.Open(44100, 16, 2)
	s = audio.NewSource(f, 1, false)
	s1 = audio.NewSource(f2, 0.4, false)
}

func (app *app) Update() {

	if !app.played {
		// fmt.Println(123)
		audio.PlaySource(s, s1)
		app.played = true
	}

}
func (app *app) Draw() {
}

func main() {

	// f, _ = os.Open("/home/arion/Downloads/12_1.wav")
	// f2, _ = limao.OpenAsset("123_1.wav")
	// p = audio.NewPlayer(44100, 16, 2)
	// audio.Open(44100, 16, 2)
	// s = audio.NewSource(f, 1, false)
	// s1 = audio.NewSource(f2, 0.4, false)

	// audio.PlaySource(s, s1)

	// f, _ = os.Open("/home/arion/Downloads/12_1.wav")
	// f2, _ = limao.OpenAsset("123_1.wav")
	// p = audio.NewPlayer(44100, 16, 2)
	// s = p.NewSource(f, 1, false)
	// s1 = p.NewSource(f2, 0.4, false)

	// p.PlaySource(s, s1)

	a := new(app)
	o := &limao.RunOpts{Width: 600, Height: 800}
	limao.Run(a, o)
	// var stop chan int
	// <-stop

}
