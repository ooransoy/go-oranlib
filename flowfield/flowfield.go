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
	"github.com/ooransoy/go-oranlib/vector"
)

var StringerOverride func(*FlowField) string
var NormalizeSet bool

type FlowField [][]vector.Vector2

func (ff *FlowField) Set(locx, locy int, v vector.Vector2) {
	if NormalizeSet {
		(*ff)[locy][locx] = v.Normalize()
	} else {
		(*ff)[locy][locx] = v
	}
}

func (ff *FlowField) Get(x, y int) vector.Vector2 {
	return (*ff)[y][x]
}

func (ff *FlowField) String() string {
	if StringerOverride != nil {
		return StringerOverride(ff)
	}
	var str string
	for _, row := range *ff {
		str += "\n"
		for _, v := range row {
			str += v.String() + ", "
		}
	}
	return str
}

func (ff *FlowField) Dimensions() vector.Vector2 {
	return vector.Vector2{
		float64(len((*ff)[0])),
		float64(len(*ff)),
	}
}

func NewFlowField(w, h int) *FlowField {
	ff := make(FlowField, h, h)
	for i, _ := range ff {
		ff[i] = make([]vector.Vector2, w, w)
	}
	return &ff
}
