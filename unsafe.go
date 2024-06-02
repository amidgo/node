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

func UnsafeAppend(contentNode, item Node) Node {
	switch contentNode.Kind() {
	case Map:
		return MapNode{
			content: append(contentNode.Content(), item),
		}
	case Array:
		return ArrayNode{
			content: append(contentNode.Content(), item),
		}
	default:
		panic("append to value node")
	}
}
