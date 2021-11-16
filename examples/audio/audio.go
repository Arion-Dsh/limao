//go:build linux || darwin || windows
// +build linux darwin windows

package main

import (
	"limao"
	"limao/audio"
	"log"
	"time"
)

func main() {
	f, err := limao.OpenAsset("2.wav")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f2, err := limao.OpenAsset("2.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	audio.Open(44100, 16, 2)
	audio.NewSource(f, 1, true)
	s := audio.NewSource(f2, 0.4, false)
	time.Sleep(2 * time.Second)
	audio.PlaySource(s)

	// p.Play()
	// time.Sleep(1 * time.Second)
	// p.Puase()
	// time.Sleep(1 * time.Second)
	// p.Resume()
	ch := make(chan int, 1)
	<-ch
}
