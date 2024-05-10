package node

import "strconv"

type BoolNode struct {
	value bool
	EmptyContentableNode
}

func MakeBoolNode(value bool) *BoolNode {
	return &BoolNode{value: value}
}

func (n *BoolNode) Kind() Kind {
	return Bool
}

func (n *BoolNode) Type() Type {
	return Value
}

func (n *BoolNode) Value() string {
	return strconv.FormatBool(n.value)
}

func (n *BoolNode) Bool() bool {
	return n.value
}
