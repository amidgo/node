package node

import (
	"strconv"
)

type FloatNode struct {
	value float64
}

func MakeFloatNode(value float64) FloatNode {
	return FloatNode{value: value}
}

func (n FloatNode) Type() Type {
	return Value
}

func (n FloatNode) Kind() Kind {
	return Float
}

func (n FloatNode) Content() []Node {
	return []Node{}
}

func (n FloatNode) Value() string {
	return strconv.FormatFloat(n.value, 'g', -1, 64)
}

func (n FloatNode) Float() float64 {
	return n.value
}
