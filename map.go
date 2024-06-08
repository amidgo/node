package node

type MapNode struct {
	content []Node
}

func MakeMapNode() MapNode {
	return MapNode{
		content: make([]Node, 0),
	}
}

func MakeMapNodeWithContent(content ...Node) MapNode {
	return MapNode{
		content: content,
	}
}

func MakeMapNodeWithSlice(content []Node) MapNode {
	return MapNode{
		content: content,
	}
}

func MakeMapNodeWithCap(capacity int) MapNode {
	return MapNode{
		content: make([]Node, 0, capacity),
	}
}

func (n MapNode) Type() Type {
	return Content
}

func (n MapNode) Kind() Kind {
	return Map
}

func (n MapNode) Content() []Node {
	return n.content
}

func (n MapNode) Value() string {
	return ""
}

func MapAppend(mapNode MapNode, key, value Node) MapNode {
	if mapNode.Kind() != Map {
		panic("map append to invalid node " + mapNode.Kind().String())
	}

	return MapNode{
		content: append(mapNode.Content(), key, value),
	}
}
