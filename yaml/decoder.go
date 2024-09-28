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
	ErrUnmarshalIntoYamlNode     = errors.New("unmarshal into yaml node")
	ErrEmptyDocumentNode         = errors.New("empty document node")
	ErrConvertMappingNodeKeyItem = errors.New("convert mapping node item")
	ErrConvertMappingNodeItem    = errors.New("convert mapping node item")
	ErrConvertSequenceNodeItem   = errors.New("convert sequence node item")
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

	iter := NewYamlNodeIterator(ynd.Content)

	for iter.HasNext() {
		key, value := iter.Next()

		keyNode, err := d.convertYamlNode(key)
		if err != nil {
			return nil, errors.Join(ErrConvertMappingNodeKeyItem, err)
		}

		contentNode, err := d.convertYamlNode(value)
		if err != nil {
			return nil, errors.Join(ErrConvertMappingNodeItem, err)
		}

		mappingNode = node.MapAppend(mappingNode, keyNode, contentNode)
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

		arrayNode = node.ArrayAppend(arrayNode, nd)
	}

	return wrapWithStyle(arrayNode, ynd.Style), nil
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
			return wrapWithStyle(node.MakeStringNode(ynd.Value), ynd.Style)
		}

		return node.MakeIntegerNode(i)
	case floatTag:
		f, err := strconv.ParseFloat(ynd.Value, 64)
		if err != nil {
			return wrapWithStyle(node.MakeStringNode(ynd.Value), ynd.Style)
		}

		return node.MakeFloatNode(f)
	default:
		return wrapWithStyle(node.MakeStringNode(ynd.Value), ynd.Style)
	}
}

func wrapWithStyle(node node.Node, style yaml.Style) node.Node {
	switch style {
	case
		yaml.TaggedStyle,
		yaml.DoubleQuotedStyle,
		yaml.SingleQuotedStyle,
		yaml.LiteralStyle,
		yaml.FoldedStyle,
		yaml.FlowStyle:
		return StyleNode(node, style)
	default:
		return node
	}
}

func StyleNode(nd node.Node, style yaml.Style) node.Node {
	return styleNode{
		node:      nd,
		yamlStyle: style,
	}
}

type styleNode struct {
	node      node.Node
	yamlStyle yaml.Style
}

func (y styleNode) Style() yaml.Style {
	return y.yamlStyle
}

func (y styleNode) Type() node.Type {
	return y.node.Type()
}

func (y styleNode) Kind() node.Kind {
	return y.node.Kind()
}

func (y styleNode) Content() []node.Node {
	return y.node.Content()
}

func (y styleNode) Value() string {
	return y.node.Value()
}
