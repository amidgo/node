package test_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/assert"
)

type MapNodeIteratorTest struct {
	CaseName        string
	MapNodeContent  []node.Node
	ExpectedContent []node.Node
}

func (c *MapNodeIteratorTest) Name() string {
	return c.CaseName
}

func (c *MapNodeIteratorTest) Test(t *testing.T) {
	content := make([]node.Node, 0, len(c.ExpectedContent))

	iterator := node.MakeMapNodeIterator(c.MapNodeContent)
	for iterator.HasNext() {
		key, value := iterator.Next()
		content = append(content, key, value)
	}

	assert.Equal(t, c.ExpectedContent, content)
}

func Test_MapNodeInterator(t *testing.T) {
	tester.RunNamedTesters(t,
		&MapNodeIteratorTest{
			CaseName:        "empty content",
			MapNodeContent:  []node.Node{},
			ExpectedContent: []node.Node{},
		},
		&MapNodeIteratorTest{
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
		&MapNodeIteratorTest{
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

type IndexedIteratorTest struct {
	CaseName      string
	Iterator      node.Iterator
	Iterations    int
	ExpectedIndex int
}

func (i *IndexedIteratorTest) Name() string {
	return i.CaseName
}

func (i *IndexedIteratorTest) Test(t *testing.T) {
	iter := node.NewIndexedIterator(i.Iterator)

	for range i.Iterations {
		iter.Next()
	}

	assert.Equal(t, i.ExpectedIndex, iter.Index())
}

func Test_IndexedIterator(t *testing.T) {
	tester.RunNamedTesters(t,
		&IndexedIteratorTest{
			CaseName:      "empty iterator",
			Iterator:      node.MakeMapNodeIterator([]node.Node{}),
			Iterations:    0,
			ExpectedIndex: -2,
		},
		&IndexedIteratorTest{
			CaseName: "iterator one iteration",
			Iterator: node.MakeMapNodeIterator([]node.Node{
				node.MakeStringNode("key"),
				node.MakeMapNode(),
			}),
			Iterations:    1,
			ExpectedIndex: 0,
		},
		&IndexedIteratorTest{
			CaseName: "iterator one iteration",
			Iterator: node.MakeMapNodeIterator([]node.Node{
				node.MakeStringNode("key"),
				node.MakeMapNode(),
			}),
			Iterations:    1,
			ExpectedIndex: 0,
		},
	)
}
