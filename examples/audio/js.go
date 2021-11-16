//go:build js
// +build js

package main

import (
	"limao"
	"limao/audio"
)

func main() {

	f, err := limao.OpenAsset("1.wav")

	p := audio.NewPlayer(44100, 16, 2)
	p.NewSource(f, 1, true)
	// p.NewSource(f2, 0.4)

	ch := make(chan int, 1)
	<-ch

}
