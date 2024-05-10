package node

import "strconv"

type IntegerNode struct {
	value int64
	EmptyContentableNode
}

func MakeIntegerNode(value int64) *IntegerNode {
	return &IntegerNode{value: value}
}

func (n *IntegerNode) Kind() Kind {
	return Integer
}

func (n *IntegerNode) Type() Type {
	return Value
}

func (n *IntegerNode) Value() string {
	return strconv.FormatInt(n.value, 10)
}

func (n *IntegerNode) Int() int64 {
	return n.value
}
