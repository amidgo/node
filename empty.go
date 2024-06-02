package node

type EmptyNode struct{}

func (n EmptyNode) Type() Type {
	return Value
}

func (n EmptyNode) Kind() Kind {
	return Empty
}

func (n EmptyNode) Content() []Node {
	return []Node{}
}

func (n EmptyNode) Value() string {
	return ""
}
