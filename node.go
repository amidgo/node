package node

type Node interface {
	Type() Type
	Kind() Kind
	Content() []Node
	Value() string
}
