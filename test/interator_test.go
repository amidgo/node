package test_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/assert"
)

func Test_MapNodeInterator(t *testing.T) {
	tester.RunNamedTesters(t,
		&MapNodeIteratorCase{
			CaseName:        "empty content",
			MapNodeContent:  []node.Node{},
			ExpectedContent: []node.Node{},
		},
		&MapNodeIteratorCase{
			CaseName: "event content length",
			MapNodeContent: []node.Node{
				node.MakeStringNode("sdkfaf"),
				node.MakeStringNode("sdksdfsdfkfaf"),

				node.MakeStringNode("adfkasdf"),
				node.MakeStringNode("sdksdasdkfglafsdfkfaf"),
			},
			ExpectedContent: []node.Node{
				node.MakeStringNode("sdkfaf"),
				node.MakeStringNode("sdksdfsdfkfaf"),

				node.MakeStringNode("adfkasdf"),
				node.MakeStringNode("sdksdasdkfglafsdfkfaf"),
			},
		},
		&MapNodeIteratorCase{
			CaseName: "odd content length",
			MapNodeContent: []node.Node{
				node.MakeStringNode("sdkfaf"),
				node.MakeStringNode("sdksdfsdfkfaf"),

				node.MakeStringNode("adfkasdf"),
			},
			ExpectedContent: []node.Node{
				node.MakeStringNode("sdkfaf"),
				node.MakeStringNode("sdksdfsdfkfaf"),

				node.MakeStringNode("adfkasdf"),
				node.EmptyNode{},
			},
		},
	)
}

type MapNodeIteratorCase struct {
	CaseName        string
	MapNodeContent  []node.Node
	ExpectedContent []node.Node
}

func (c *MapNodeIteratorCase) Name() string {
	return c.CaseName
}

func (c *MapNodeIteratorCase) Test(t *testing.T) {
	content := make([]node.Node, 0, len(c.ExpectedContent))

	iterator := node.MakeMapNodeIterator(c.MapNodeContent)
	for iterator.HasNext() {
		key, value := iterator.Next()
		content = append(content, key, value)
	}

	assert.Equal(t, c.ExpectedContent, content)
}
