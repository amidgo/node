package node

type ArrayNode struct {
	BaseContentableNode
}

func MakeArrayNode() *ArrayNode {
	return &ArrayNode{
		BaseContentableNode: BaseContentableNode{
			content: make([]Node, 0),
		},
	}
}

func MakeArrayNodeWithContent(content ...Node) *ArrayNode {
	return &ArrayNode{
		BaseContentableNode: BaseContentableNode{
			content: content,
		},
	}
}

func (h *ArrayNode) Type() Type {
	return Array
}
