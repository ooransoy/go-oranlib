/*
	Copyright (C) 2018 Olcay Oransoy

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Package vector provides tools to simplify vector operations. In the current
// state, it only supports 2D vectors. See the code and its comments to learn
// more about how to use this package.
package vector

import (
	"fmt"
	"math"
)

// If this function isn't nil, the ``(v Vector2) String() string'' function
// uses this to stringify vectors.
var StringerOverride func(Vector2) string

// Vector2 is a two-dimensional vector. Pretty standard stuff.
type Vector2 struct {
	X float64
	Y float64
}

// Normalize takes a Vector2 as the reciever and turns it into a unit vector.
// What this means is that the resulting vector's magnitude is 1.
// The square root of the sum of a unit vector's components' squares is equal
// to one. sqrt(pow(x, 2) + pow(y, 2)) == 1
// Aww, no pointless "documentation" "jokes" here?
func (v Vector2) Normalize() Vector2 {
	magnitude := math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
	return Vector2{v.X / magnitude, v.Y / magnitude}
}

// Scale scales the Vector2 given as the reciever by the float given in the
// arguments. Why do you even need documentation for this? It is a helper
// function consisting of a single line.
func (v Vector2) Scale(s float64) Vector2 {
	return Vector2{v.X * s, v.Y * s}
}

// String stringifies the Vector2 reciever.
// Actually, if M-Theory is real it doesn't have to.
func (v Vector2) String() string {
	if StringerOverride != nil {
		return StringerOverride(v)
	}
	return fmt.Sprintf("(%07.3f, %07.3f)", v.X, v.Y)
}
