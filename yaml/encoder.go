package yaml

import (
	"bytes"
	"errors"

	"github.com/amidgo/node"
	"gopkg.in/yaml.v3"
)

var (
	ErrConvertToYamlNode    = errors.New("convert to yaml node")
	ErrMarshalYamlNode      = errors.New("marshal yaml node")
	ErrConvertMapNodeItem   = errors.New("convert map node item")
	ErrConvertArrayNodeItem = errors.New("convert array node item")
)

type Encoder struct {
	Indent int
}

func (e *Encoder) Encode(nd node.Node) ([]byte, error) {
	ynd, err := e.convertNode(nd)
	if err != nil {
		return nil, errors.Join(ErrConvertToYamlNode, err)
	}

	ynd = &yaml.Node{
		Kind:    yaml.DocumentNode,
		Content: []*yaml.Node{ynd},
	}

	out := bytes.Buffer{}

	enc := yaml.NewEncoder(&out)
	defer enc.Close()

	enc.SetIndent(e.Indent)

	err = enc.Encode(ynd)
	if err != nil {
		return nil, errors.Join(ErrMarshalYamlNode, err)
	}

	return out.Bytes(), nil
}

func Encode(nd node.Node) ([]byte, error) {
	enc := new(Encoder)

	return enc.Encode(nd)
}

func EncodeWithIndent(nd node.Node, indent int) ([]byte, error) {
	enc := new(Encoder)

	enc.Indent = indent

	return enc.Encode(nd)
}

func (e *Encoder) convertNode(nd node.Node) (*yaml.Node, error) {
	switch nd.Type() {
	case node.Map:
		return e.convertMapNode(nd)
	case node.Array:
		return e.convertArrayNode(nd)
	case node.Value:
		return e.convertValueNode(nd)
	default:
		return nil, &UnsupportedNodeTypeError{InputType: nd.Type()}
	}
}

func (e *Encoder) convertMapNode(nd node.Node) (*yaml.Node, error) {
	yndContent := make([]*yaml.Node, 0, len(nd.Content()))

	for _, item := range nd.Content() {
		ynd, err := e.convertNode(item)
		if err != nil {
			return nil, errors.Join(ErrConvertMappingNodeItem)
		}

		yndContent = append(yndContent, ynd)
	}

	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Content: yndContent,
	}, nil
}

func (e *Encoder) convertArrayNode(nd node.Node) (*yaml.Node, error) {
	yndContent := make([]*yaml.Node, 0, len(nd.Content()))

	for _, item := range nd.Content() {
		ynd, err := e.convertNode(item)
		if err != nil {
			return nil, errors.Join(ErrConvertArrayNodeItem)
		}

		yndContent = append(yndContent, ynd)
	}

	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Content: yndContent,
		Style:   extractYamlStyleFromNode(nd),
	}, nil
}

func (e *Encoder) convertValueNode(nd node.Node) (*yaml.Node, error) {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: nd.Value(),
		Style: extractYamlStyleFromNode(nd),
	}, nil
}

func extractYamlStyleFromNode(nd node.Node) yaml.Style {
	var style yaml.Style

	styledNode, ok := nd.(interface{ Style() yaml.Style })
	if ok {
		style = styledNode.Style()
	}

	return style
}
