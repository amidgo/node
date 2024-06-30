package node

type StringNode struct {
	value string
}

func MakeStringNode(value string) StringNode {
	return StringNode{
		value: value,
	}
}

func (s StringNode) Type() Type {
	return Value
}

func (s StringNode) Kind() Kind {
	return String
}

func (s StringNode) Content() []Node {
	return []Node{}
}

func (s StringNode) Value() string {
	return s.value
}

func StringEquals(nd Node, s string) bool {
	return nd.Kind() == String && nd.Value() == s
}
