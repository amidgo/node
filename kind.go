package node

type Kind int

const (
	String Kind = iota
	Integer
	Float
	Bool
	Empty
	Content
)
