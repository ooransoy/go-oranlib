/*
	FYI: This code was written with the intention of manually testing the
	ffagent library.
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/ooransoy/go-oranlib/ffagent"
	flow "github.com/ooransoy/go-oranlib/flowfield"
	"github.com/ooransoy/go-oranlib/vector"
)

// RNG Seeder
var _ = func() struct{} { rand.Seed(time.Now().UnixNano()); return struct{}{} }()

const (
	width  = 800
	height = 600
)

func init() {
	flow.NormalizeSet = true
	ffagent.WrapEdges = true
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	a := ffagent.Agent{}
	a.Field = flow.NewFlowField(width, height)
	a.Location = vector.Vector2{
		rand.Float64() * width,
		rand.Float64() * height,
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			a.Field.Set(
				x,
				y,
				vector.Vector2{
					float64(width-x) + rand.Float64()*500,
					float64(height/2.0-y) + rand.Float64()*500,
				},
			)
		}
	}

	for i := 0; i < 1000; i++ {
		a.Advance()
		img.Set(
			int(a.Location.X),
			int(a.Location.Y),
			color.RGBA{255, 0, 255, 255},
		)
	}

	if err := save(img, "images/img.png"); err != nil {
		panic(err)
	}
}

func save(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
