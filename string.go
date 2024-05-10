package node

type StringNode struct {
	value string
	EmptyContentableNode
}

func MakeStringNode(value string) *StringNode {
	return &StringNode{
		value: value,
	}
}

func (s *StringNode) Value() string {
	return s.value
}

func (s *StringNode) Kind() Kind {
	return String
}

func (s *StringNode) Type() Type {
	return Value
}
