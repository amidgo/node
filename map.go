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
	return MapNode{
		content: append(mapNode.Content(), key, value),
	}
}

func MapSet(mapNode MapNode, key, value Node) MapNode {
	iter := NewIndexedIterator(MakeMapNodeIterator(mapNode.Content()))

	for iter.HasNext() {
		nextKey, _ := iter.Next()
		if nextKey != key {
			continue
		}

		if iter.Index() == len(mapNode.Content())-1 {
			return MapNode{
				content: append(mapNode.Content(), value),
			}
		}

		mapNode.Content()[iter.Index()+1] = value

		return mapNode
	}

	return MapAppend(mapNode, key, value)
}

// if len(mapNode.Content()) % 2 == 1, append value node to Map and return them
func MapRound(mapNode MapNode, value Node) MapNode {
	if len(mapNode.Content())%2 == 0 {
		return mapNode
	}

	return MapNode{
		content: append(mapNode.Content(), value),
	}
}
