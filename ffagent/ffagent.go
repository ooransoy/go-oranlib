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
