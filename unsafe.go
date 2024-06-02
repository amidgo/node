package node

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

func (n *UnsafeNode) Value() string {
	return n.NValue
}

func Append(contentNode, item Node) {
	switch contentNode := contentNode.(type) {
	case *MapNode:
		contentNode.content = append(contentNode.content, item)
	case *ArrayNode:
		contentNode.content = append(contentNode.content, item)
	case *UnsafeNode:
		contentNode.NContent = append(contentNode.NContent, item)
	}
}
