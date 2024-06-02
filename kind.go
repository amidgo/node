package node

type Kind int

const (
	String Kind = iota
	Integer
	Float
	Bool
	Empty
	Map
	Array
)

func (k Kind) String() string {
	switch k {
	case String:
		return "String"
	case Integer:
		return "Integer"
	case Float:
		return "Float"
	case Bool:
		return "Bool"
	case Empty:
		return "Empty"
	case Map:
		return "Map"
	case Array:
		return "Content"
	default:
		return "Unknown Kind"
	}
}
