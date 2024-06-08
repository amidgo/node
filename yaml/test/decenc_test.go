package test_test

import (
	_ "embed"
	"testing"

	"github.com/amidgo/node/yaml"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata/swagger.yaml
	swaggerData []byte
	//go:embed testdata/simple.yaml
	simpleSwaggerData []byte
	//go:embed testdata/simple_with_array.yaml
	simpleSwaggerWithArrayData []byte
)

type DecEncTester struct {
	CaseName string
	Indent   int
	Input    []byte
}

func (de *DecEncTester) Name() string {
	return "decenc tester " + de.CaseName
}

func (de *DecEncTester) Test(t *testing.T) {
	nd, err := yaml.Decode(de.Input)
	require.NoError(t, err)

	data, err := yaml.EncodeWithIndent(nd, de.Indent)
	require.NoError(t, err)

	t.Logf("encodedLen:%d inputLen:%d", len(data), len(de.Input))
	t.Logf("encoded:\n%s\ninput:\n%s", string(data), string(de.Input))

	assert.Equal(t, string(de.Input), string(data))
}

func Test_DecEnc(t *testing.T) {
	tester.RunNamedTesters(t,
		&DecEncTester{
			CaseName: "simple swagger 3.0.1",
			Indent:   2,
			Input:    simpleSwaggerData,
		},
		&DecEncTester{
			CaseName: "simple swagger with array 3.0.1",
			Indent:   2,
			Input:    simpleSwaggerWithArrayData,
		},
		&DecEncTester{
			CaseName: "full swagger 3.0.1",
			Indent:   2,
			Input:    swaggerData,
		},
	)
}
