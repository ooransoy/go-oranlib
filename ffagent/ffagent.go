package ffagent

import (
	"math"

	flow "github.com/ooransoy/go-oranlib/flowfield"
	"github.com/ooransoy/go-oranlib/vector"
)

type Agent struct {
	Field    *flow.FlowField
	Location vector.Vector2
}

var WrapEdges = true

func (a *Agent) Advance() {
	standingOn := a.Field.Get(int(math.Floor(a.Location.X)), int(math.Floor(a.Location.Y)))
	a.Location.X += standingOn.X
	a.Location.Y += standingOn.Y

	if WrapEdges {
		d := a.Field.Dimensions()

	xCheck:
		if a.Location.X >= d.X {
			a.Location.X -= d.X
			goto xCheck
		}

	yCheck:
		if a.Location.Y >= d.Y {
			a.Location.Y -= d.Y
			goto yCheck
		}
	}
}
