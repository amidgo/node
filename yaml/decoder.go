package yaml

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/amidgo/node"
	"gopkg.in/yaml.v3"
)

var (
	ErrUnmarshalIntoYamlNode   = errors.New("unmarshal into yaml node")
	ErrEmptyDocumentNode       = errors.New("empty document node")
	ErrConvertMappingNodeItem  = errors.New("convert mapping node item")
	ErrConvertSequenceNodeItem = errors.New("convert sequence node item")
)

type Decoder struct{}

func (d *Decoder) Decode(data []byte) (node.Node, error) {
	var yamlNode yaml.Node

	err := yaml.Unmarshal(data, &yamlNode)
	if err != nil {
		return nil, errors.Join(ErrUnmarshalIntoYamlNode, err)
	}

	return d.convertYamlNode(&yamlNode)
}

func (d *Decoder) DecodeFrom(r io.Reader) (node.Node, error) {
	var yamlNode yaml.Node

	dec := yaml.NewDecoder(r)

	err := dec.Decode(&yamlNode)
	if err != nil {
		return nil, errors.Join(ErrUnmarshalIntoYamlNode, err)
	}

	return d.convertYamlNode(&yamlNode)
}

func Decode(data []byte) (node.Node, error) {
	dec := new(Decoder)

	return dec.Decode(data)
}

func (d *Decoder) convertYamlNode(ynd *yaml.Node) (node.Node, error) {
	switch ynd.Kind {
	case yaml.DocumentNode:
		if len(ynd.Content) != 1 {
			return nil, ErrEmptyDocumentNode
		}

		return d.convertYamlNode(ynd.Content[0])
	case yaml.MappingNode:
		return d.convertMappingYamlNode(ynd)
	case yaml.SequenceNode:
		return d.convertSequenceYamlNode(ynd)
	case yaml.ScalarNode:
		return convertScalarYamlNode(ynd), nil
	default:
		return nil, &UnsupportedYamlNodeKindError{InputKind: ynd.Kind}
	}
}

func (d *Decoder) convertMappingYamlNode(ynd *yaml.Node) (node.Node, error) {
	mappingNode := node.MakeMapNode()

	for _, item := range ynd.Content {
		nd, err := d.convertYamlNode(item)
		if err != nil {
			return nil, errors.Join(ErrConvertMappingNodeItem, err)
		}

		mappingNode.AppendNode(nd)
	}

	return mappingNode, nil
}

func (d *Decoder) convertSequenceYamlNode(ynd *yaml.Node) (node.Node, error) {
	arrayNode := node.MakeArrayNode()

	for _, item := range ynd.Content {
		nd, err := d.convertYamlNode(item)
		if err != nil {
			return nil, errors.Join(ErrConvertSequenceNodeItem, err)
		}

		arrayNode.AppendNode(nd)
	}

	return &YamlStyleNode{Node: arrayNode, YamlStyle: ynd.Style}, nil
}

const (
	nullTag  = "!!null"
	boolTag  = "!!bool"
	intTag   = "!!int"
	floatTag = "!!float"
)

func convertScalarYamlNode(ynd *yaml.Node) node.Node {
	switch ynd.Tag {
	case nullTag:
		return node.EmptyNode{}
	case boolTag:
		return node.MakeBoolNode(strings.ToLower(ynd.Value) == "true")
	case intTag:
		i, err := strconv.ParseInt(ynd.Value, 10, 64)
		if err != nil {
			return stringNode(ynd.Value, ynd.Style)
		}

		return node.MakeIntegerNode(i)
	case floatTag:
		f, err := strconv.ParseFloat(ynd.Value, 64)
		if err != nil {
			return stringNode(ynd.Value, ynd.Style)
		}

		return node.MakeFloatNode(f)
	default:
		return stringNode(ynd.Value, ynd.Style)
	}
}

func stringNode(value string, style yaml.Style) node.Node {
	return &YamlStyleNode{Node: node.MakeStringNode(value), YamlStyle: style}
}

type YamlStyleNode struct {
	node.Node
	YamlStyle yaml.Style
}

func (y *YamlStyleNode) Style() yaml.Style {
	return y.YamlStyle
}
