package node

import (
	"bytes"
	"fmt"
	"sync"
)

type Kind int

const (
	Empty Kind = iota
	Integer
	Float
	Bool
	String
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
		return "Array"
	default:
		return "Unknown Kind"
	}
}

type InvalidNodeKindError struct {
	ActualKind Kind
}

func (i InvalidNodeKindError) Error() string {
	return "invalid node kind, actual " + i.ActualKind.String()
}

func (i InvalidNodeKindError) Is(err error) bool {
	invalidNodeKindError, ok := err.(InvalidNodeKindError)
	if !ok {
		return false
	}

	return invalidNodeKindError.ActualKind == i.ActualKind
}

type KindValidate struct {
	validKinds []Kind
	kinds      string
	once       sync.Once
}

func NewKindValidate(validKinds ...Kind) *KindValidate {
	return &KindValidate{
		validKinds: validKinds,
	}
}

func (n *KindValidate) Validate(nd Node) error {
	for _, kind := range n.validKinds {
		if kind == nd.Kind() {
			return nil
		}
	}

	n.once.Do(n.setKinds)

	return fmt.Errorf("%w, expected one of [%s]", InvalidNodeKindError{ActualKind: nd.Kind()}, n.kinds)
}

func (n *KindValidate) setKinds() {
	kinds := &bytes.Buffer{}

	for index, kind := range n.validKinds {
		kinds.WriteString(kind.String())

		if index != len(n.validKinds)-1 {
			kinds.WriteByte(',')
			kinds.WriteByte(' ')
		}
	}

	n.kinds = kinds.String()
}
