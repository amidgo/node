package node

type MapNode struct {
	BaseContentableNode
}

func MakeMapNode() *MapNode {
	return &MapNode{
		BaseContentableNode: BaseContentableNode{
			content: make([]Node, 0),
		},
	}
}

func MakeMapNodeWithContent(content ...Node) *MapNode {
	return &MapNode{
		BaseContentableNode: BaseContentableNode{
			content: content,
		},
	}
}

func MakeMapNodeWithCap(cap int) *MapNode {
	return &MapNode{
		BaseContentableNode: BaseContentableNode{
			content: make([]Node, 0, cap),
		},
	}
}

func (h *MapNode) Type() Type {
	return Map
}
