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

package flowfield

import (
	"testing"

	vec "github.com/ooransoy/go-oranlib/vector"
)

var coords = [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}

func TestGet(t *testing.T) {
	ff := &FlowField{
		{vec.Vector2{0, 0}, vec.Vector2{1, 0}},
		{vec.Vector2{0, 1}, vec.Vector2{1, 1}},
	}

	for _, coord := range coords {
		v := ff.Get(coord[0], coord[1])
		expected := fromIntArr(coord)
		if v != expected {
			t.Errorf("expected %v but got %v", expected, v)
		}
	}
}

func TestSet(t *testing.T) {
	testVector := vec.Vector2{4, 3}

	for _, coord := range coords {
		ff := NewFlowField(2, 2)
		ff.Set(coord[0], coord[1], testVector)
		actual := ff.Get(coord[0], coord[1])
		if actual != testVector {
			t.Errorf("expected %v but got %v", testVector, actual)
		}
	}

	NormalizeSet = true

	expected := testVector.Normalize()
	for _, coord := range coords {
		ff := NewFlowField(2, 2)
		ff.Set(coord[0], coord[1], testVector)
		actual := ff.Get(coord[0], coord[1])
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	}
}

func TestDimensions(t *testing.T) {
	w := 800.0
	h := 600.0
	d := NewFlowField(int(w), int(h)).Dimensions()

	if d.X != w {
		t.Errorf("expected %v but got %v", w, d.X)
	}
	if d.Y != h {
		t.Errorf("expected %v but got %v", h, d.Y)
	}
}

func fromIntArr(arr [2]int) vec.Vector2 {
	return vec.Vector2{float64(arr[0]), float64(arr[1])}
}
