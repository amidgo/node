package node

type ArrayNode struct {
	content []Node
}

func MakeArrayNode() ArrayNode {
	return ArrayNode{
		content: make([]Node, 0),
	}
}

func MakeArrayNodeWithContent(content ...Node) ArrayNode {
	return ArrayNode{
		content: content,
	}
}

func MakeArrayNodeWithCap(cap int) ArrayNode {
	return ArrayNode{
		content: make([]Node, 0, cap),
	}
}

func (h ArrayNode) Type() Type {
	return Content
}

func (n ArrayNode) Kind() Kind {
	return Array
}

func (n ArrayNode) Content() []Node {
	return n.content
}

func (n ArrayNode) Value() string {
	return ""
}

func ArrayAppend(arrayNode Node, items ...Node) Node {
	if arrayNode.Kind() != Array {
		panic("map append to invalid node " + arrayNode.Kind().String())
	}

	return ArrayNode{
		content: append(arrayNode.Content(), items...),
	}
}
