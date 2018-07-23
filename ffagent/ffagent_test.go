package ffagent

import (
	"testing"

	flow "github.com/ooransoy/go-oranlib/flowfield"
	"github.com/ooransoy/go-oranlib/vector"
)

func TestAdvance(t *testing.T) {
	w := WrapEdges
	WrapEdges = false
	a := &Agent{}
	a.Field = flow.NewFlowField(2, 2)
	a.Field.Set(0, 0, vector.Vector2{1, 1})
	a.Advance()

	expected := vector.Vector2{1, 1}
	if a.Location != expected {
		t.Errorf("expected %v but got %v", expected, a.Location)
	}

	WrapEdges = true
	a.Field.Set(1, 1, vector.Vector2{101, 100})
	a.Advance()

	expected = vector.Vector2{0, 1}
	if a.Location != expected {
		t.Errorf("expected %v but got %v", expected, a.Location)
	}
	WrapEdges = w
}
