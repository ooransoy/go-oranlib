/*
	FYI: This code was written with the intention of manually testing the
	ffagent library.

	This file is covered by the Do What the Fuck You Want To Public License.
	Not gee enn yuu something something v3.
*/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
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
}

func main() {
	fmt.Println("Starting")
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	fmt.Println("Creating agents")
	agents := make([]*ffagent.Agent, 1000)
	for i, _ := range agents {
		agents[i] = createAgent()
	}

	fmt.Println("Initializing flow fields")
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for i, a := range agents {
				xComp := float64(width)
				yComp := float64(height)
				switch i % 5 {
				case 0:
					xComp /= 2
				case 1:
					yComp /= 2
				case 2:
					xComp += width
				case 3:
					yComp += height
				}
				xComp -= float64(x)
				yComp -= float64(y)
				a.Field.Set(
					int(x),
					int(y),
					vector.Vector2{
						xComp + rand.Float64()*1,
						yComp + rand.Float64()*1,
					},
				)
			}
		}
	}

	fmt.Println("Advancing agents")
	for i := 0; i < 250; i++ {
		for _, a := range agents {
			a.Advance()
			xRatio := a.Location.X / width
			yRatio := a.Location.Y / height
			img.Set(
				int(a.Location.X),
				int(a.Location.Y),
				color.RGBA{
					uint8(math.Floor(xRatio * 255)), 0,
					uint8(math.Floor(yRatio * 255)),
					255,
				},
			)
		}
	}

	fmt.Println("Saving results")
	if err := save(img, "images/img.png"); err != nil {
		panic(err)
	}

	fmt.Println("Saving flowfields")
	for i, a := range agents {
		if i == 5 {
			break
		}
		ffimg := image.NewRGBA(image.Rect(0, 0, width, height))
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				v := a.Field.Get(x, y)
				ffimg.Set(
					x,
					y,
					color.RGBA{uint8(v.X*256-1), 0, uint8(v.Y*256-1), 255},
				)
			}
		}
		if err := save(
			ffimg,
			fmt.Sprintf("images/agent%d.png", i),
		); err != nil {
			panic(err)
		}
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

func createAgent() *ffagent.Agent {
	a := ffagent.Agent{}
	a.Field = flow.NewFlowField(width, height)
	a.Location = vector.Vector2{
		rand.Float64() * width,
		rand.Float64() * height,
	}
	return &a
}
