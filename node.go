package node

type Node interface {
	Type() Type
	Contentable
	Valuable
}

type UnsafeNode struct {
	NType    Type
	NKind    Kind
	NValue   string
	NContent []Node
}

func (n *UnsafeNode) Type() Type {
	return n.NType
}

func (n *UnsafeNode) Kind() Kind {
	return n.NKind
}

func (n *UnsafeNode) Content() []Node {
	return n.NContent
}

func (n *UnsafeNode) SetContent(content []Node) {
	n.NContent = content
}

func (n *UnsafeNode) AppendNode(node Node) {
	n.NContent = append(n.NContent, node)
}

func (n *UnsafeNode) Value() string {
	return n.NValue
}
