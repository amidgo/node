package yaml

import (
	"gopkg.in/yaml.v3"
)

type NodeIterator struct {
	nodes []*yaml.Node
}

func NewYamlNodeIterator(content []*yaml.Node) NodeIterator {
	return NodeIterator{
		nodes: content,
	}
}

func (i NodeIterator) HasNext() bool {
	return len(i.nodes) >= 2
}

func (i *NodeIterator) Next() (key, value *yaml.Node) {
	key, value = i.nodes[0], i.nodes[1]

	if len(i.nodes) >= 2 {
		i.nodes = i.nodes[2:]
	}

	return key, value
}
