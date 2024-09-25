package node

type Iterator interface {
	HasNext() bool
	Next() (key, value Node)
}

type mapNodeIterator struct {
	content []Node
}

func MakeMapNodeIterator(content []Node) Iterator {
	return &mapNodeIterator{
		content: content,
	}
}

func (i *mapNodeIterator) HasNext() bool {
	return len(i.content) >= 1
}

func (i *mapNodeIterator) Next() (key, value Node) {
	key = i.content[0]

	if len(i.content) == 1 {
		i.content = i.content[:0]

		return key, EmptyNode{}
	}

	value = i.content[1]
	i.content = i.content[2:]

	return key, value
}

type IndexedIterator struct {
	iter     Iterator
	keyIndex int
}

func NewIndexedIterator(iter Iterator) *IndexedIterator {
	return &IndexedIterator{
		iter:     iter,
		keyIndex: -2,
	}
}

func (i *IndexedIterator) HasNext() bool {
	return i.iter.HasNext()
}

func (i *IndexedIterator) Next() (key, value Node) {
	i.keyIndex += 2

	return i.iter.Next()
}

func (i *IndexedIterator) Index() int {
	return i.keyIndex
}
