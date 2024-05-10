package node

type Type int

const (
	Map Type = iota
	Array
	Value
)

func (t Type) String() string {
	switch t {
	case Map:
		return "Map"
	case Array:
		return "Array"
	case Value:
		return "Value"
	default:
		return ""
	}
}
