package node

type Valuable interface {
	Value() string
	Kind() Kind
}

type EmptyValueableNode struct{}

func (n EmptyValueableNode) Value() string {
	return ""
}
