package node

type MapNodeIterator struct {
	content []Node
	offset  int
}

func MakeMapNodeIterator(content []Node) *MapNodeIterator {
	return &MapNodeIterator{
		content: content,
	}
}

func (i *MapNodeIterator) HasNext() bool {
	return i.offset < len(i.content)
}

func (i *MapNodeIterator) Next() (key, value Node) {
	key = i.content[i.offset]

	if i.offset+1 < len(i.content) {
		value = i.content[i.offset+1]
	} else {
		value = EmptyNode{}
	}

	i.offset += 2

	return key, value
}
