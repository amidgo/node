package node

type Contentable interface {
	Content() []Node
	SetContent(content []Node)
	AppendNode(node Node)
}

type BaseContentableNode struct {
	content []Node
	EmptyValueableNode
}

func (n *BaseContentableNode) Content() []Node {
	return n.content
}

func (n *BaseContentableNode) SetContent(content []Node) {
	n.content = content
}

func (n *BaseContentableNode) AppendNode(node Node) {
	n.content = append(n.content, node)
}

func (n *BaseContentableNode) Kind() Kind {
	return Content
}

type EmptyContentableNode struct{}

func (n EmptyContentableNode) Content() []Node {
	return []Node{}
}

func (n EmptyContentableNode) SetContent(_ []Node) {
}

func (n EmptyContentableNode) AppendNode(_ Node) {
}
