package node

type LineableNode interface {
	Node
	Line() int
	Column() int
}

type lineableNode struct {
	line, column int
	Node
}

func (l *lineableNode) Line() int {
	return l.line
}

func (l *lineableNode) Column() int {
	return l.column
}

func MakeLineableNode(nd Node, line, column int) LineableNode {
	return &lineableNode{Node: nd, line: line, column: column}
}
