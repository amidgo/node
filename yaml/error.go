package yaml

import (
	"fmt"

	"github.com/amidgo/node"
	"gopkg.in/yaml.v3"
)

type UnsupportedYamlNodeKindError struct {
	InputKind yaml.Kind
}

func (e *UnsupportedYamlNodeKindError) Error() string {
	return fmt.Sprintf("unsupported node kind: %d", e.InputKind)
}

func (e *UnsupportedYamlNodeKindError) Is(target error) bool {
	unsupportKindErr, ok := target.(*UnsupportedYamlNodeKindError)
	if ok {
		return unsupportKindErr.InputKind == e.InputKind
	}

	return false
}

type UnsupportedNodeTypeError struct {
	InputType node.Type
}

func (e *UnsupportedNodeTypeError) Error() string {
	return fmt.Sprintf("unsupported node type: %s", e.InputType)
}

func (e *UnsupportedNodeTypeError) Is(target error) bool {
	unsupportedNodeTypeErr, ok := target.(*UnsupportedNodeTypeError)
	if ok {
		return unsupportedNodeTypeErr.InputType == e.InputType
	}

	return false
}
