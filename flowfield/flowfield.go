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
