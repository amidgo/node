package node

type ArrayNode struct {
	content []Node
}

func MakeArrayNode(content ...Node) ArrayNode {
	return ArrayNode{
		content: content,
	}
}

func (n ArrayNode) Type() Type {
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

func ArrayAppend(arrayNode ArrayNode, items ...Node) ArrayNode {
	return ArrayNode{
		content: append(arrayNode.Content(), items...),
	}
}
