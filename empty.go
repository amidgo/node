package node

type EmptyNode struct {
	EmptyContentableNode
}

func (n EmptyNode) Kind() Kind {
	return Empty
}

func (n EmptyNode) Type() Type {
	return Value
}

func (n EmptyNode) Value() string {
	return ""
}
