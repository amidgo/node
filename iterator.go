package node

type MapNodeIterator struct {
	content []Node
}

func MakeMapNodeIterator(content []Node) *MapNodeIterator {
	return &MapNodeIterator{
		content: content,
	}
}

func (i *MapNodeIterator) HasNext() bool {
	return len(i.content) >= 1
}

func (i *MapNodeIterator) Next() (key, value Node) {
	key = i.content[0]

	if len(i.content) == 1 {
		i.content = i.content[:0]

		return key, EmptyNode{}
	}

	value = i.content[1]
	i.content = i.content[2:]

	return key, value
}
