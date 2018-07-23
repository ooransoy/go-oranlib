package vector

import (
	"fmt"
	"math"
)

var StringerOverride func(Vector2) string

// Vector2 is a two-dimensional vector. Pretty standard stuff.
type Vector2 struct {
	X float64
	Y float64
}

// Normalize takes a Vector2 and turns it into a unit vector.
// What this means is that the resulting vector's magnitude is 1.
// The square root of the sum of a unit vector's components' squares is equal
// to one. sqrt(pow(x, 2) + pow(y, 2)) == 1
func (v Vector2) Normalize() Vector2 {
	magnitude := math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
	return Vector2{v.X / magnitude, v.Y / magnitude}
}

func Scale(v Vector2, s float64) Vector2 {
	return Vector2{v.X * s, v.Y * s}
}

func (v Vector2) String() string {
	if StringerOverride != nil {
		return StringerOverride(v)
	}
	return fmt.Sprintf("(%07.3f, %07.3f)", v.X, v.Y)
}
