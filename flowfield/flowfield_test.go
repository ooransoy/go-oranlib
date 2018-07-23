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
