package jsontest

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type DecodeTestCase struct {
	CaseName string

	Data         string
	ExpectedNode node.Node
	ExpectedErr  error
}

func (c *DecodeTestCase) Name() string {
	return c.CaseName
}

func (c *DecodeTestCase) Test(t *testing.T) {
	decoder := new(json.Decoder)

	initNode, err := decoder.Decode([]byte(c.Data))

	require.ErrorIs(t, err, c.ExpectedErr)

	assert.Equal(t, c.ExpectedNode, initNode)
}

type EncodeTestCase struct {
	CaseName string

	Node         node.Node
	ExpectedData string
	ExpectedErr  error
}

func (c *EncodeTestCase) Name() string {
	return c.CaseName
}

func (c *EncodeTestCase) Test(t *testing.T) {
	encoder := new(json.Encoder)

	data, err := encoder.Encode(c.Node)

	require.ErrorIs(t, err, c.ExpectedErr)

	assert.Equal(t, c.ExpectedData, string(data))
}
