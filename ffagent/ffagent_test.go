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
