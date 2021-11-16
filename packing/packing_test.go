package packing

import (
	"image"
	"log"
	"testing"
)

func TestNode(t *testing.T) {

	root := &node{x: 0, y: 0, w: 300, h: 200, left: new(node), right: new(node)}

	rect := &node{w: 100, h: 100}
	rect2 := &node{w: 20, h: 20}
	rect3 := &node{w: 120, h: 120}

	a := root.insert(rect, true)
	if a.x != 0 || a.y != 0 {
		t.Fail()
	}
	b := root.insert(rect2, true)
	if b.x != 0 || b.y != 100 {
		t.Fail()
	}
	c := root.insert(rect3, true)
	if c.x != 100 || c.y != 0 {
		t.Fail()
	}

}

func TestNodeFull(t *testing.T) {

	root := &node{x: 0, y: 0, w: 300, h: 200, left: new(node), right: new(node)}

	rect := &node{w: 300, h: 300}

	a := root.insert(rect, true)

	if a != nil {
		t.Fail()

	}

}

func TestPacker(t *testing.T) {
	p := New(100, 100)

	rect := image.Rect(0, 0, 300, 300)
	r, err := p.Insert(rect)

	if (r != image.Rectangle{}) || err == nil {
		t.Fail()
	}

	if err != nil {
		p.Grow(400, 400)
	}

	r, err = p.Insert(rect)

	if r != image.Rect(0, 100, 300, 400) || err != nil {
		log.Fatal()
	}

	r, err = p.Insert(image.Rect(0, 0, 50, 50))

	if r != image.Rect(0, 0, 50, 50) || err != nil {
		log.Fatal()
	}

}
