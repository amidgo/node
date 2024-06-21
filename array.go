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

func MakeArrayNodeWithCap(capacity int) ArrayNode {
	return ArrayNode{
		content: make([]Node, 0, capacity),
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
