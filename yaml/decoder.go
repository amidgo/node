package yaml

import (
	"errors"
	"io"

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
		return d.convertScalarYamlNode(ynd)
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

	return &yamlStyleNode{Node: arrayNode, style: ynd.Style}, nil
}

func (d *Decoder) convertScalarYamlNode(ynd *yaml.Node) (node.Node, error) {
	return &yamlStyleNode{Node: node.MakeStringNode(ynd.Value), style: ynd.Style}, nil
}

type yamlStyleNode struct {
	node.Node
	style yaml.Style
}

func (y *yamlStyleNode) Style() yaml.Style {
	return y.style
}
