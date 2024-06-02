package yaml

import (
	"gopkg.in/yaml.v3"
)

type YamlNodeIterator struct {
	nodes []*yaml.Node
}

func NewYamlNodeIterator(content []*yaml.Node) YamlNodeIterator {
	return YamlNodeIterator{
		nodes: content,
	}
}

func (i YamlNodeIterator) HasNext() bool {
	return len(i.nodes) >= 2
}

func (i *YamlNodeIterator) Next() (key, value *yaml.Node) {
	key, value = i.nodes[0], i.nodes[1]

	if len(i.nodes) >= 2 {
		i.nodes = i.nodes[2:]
	}

	return key, value
}
