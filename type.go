package node

type Type int

const (
	Content Type = iota + 1
	Value
)

func (t Type) String() string {
	switch t {
	case Content:
		return "Content"
	case Value:
		return "Value"
	default:
		return ""
	}
}
