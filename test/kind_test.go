package test_test

import (
	"strings"
	"sync"
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/require"
)

type KindValidateTest struct {
	CaseName   string
	ValidKinds []node.Kind
	Node       node.Node

	ExpectedError error
}

func (k *KindValidateTest) Name() string {
	return k.CaseName
}

func (k *KindValidateTest) Test(t *testing.T) {
	validate := node.NewKindValidate(k.ValidKinds...)

	err := validate.Validate(k.Node)
	require.ErrorIs(t, err, k.ExpectedError)
}

func Test_KindValidate(t *testing.T) {
	tester.RunNamedTesters(t,
		&KindValidateTest{
			CaseName:      "empty kinds test",
			ValidKinds:    nil,
			Node:          node.MakeMapNode(),
			ExpectedError: node.InvalidNodeKindError{ActualKind: node.Map},
		},
		&KindValidateTest{
			CaseName:      "valid kinds match with node kind",
			ValidKinds:    []node.Kind{node.Array, node.Bool, node.String},
			Node:          node.MakeStringNode("Hello World!"),
			ExpectedError: nil,
		},
		&KindValidateTest{
			CaseName:      "valid kinds not match with one kind",
			ValidKinds:    []node.Kind{node.Array, node.Bool, node.String},
			Node:          node.MakeMapNode(),
			ExpectedError: node.InvalidNodeKindError{ActualKind: node.Map},
		},
	)
}

func Test_KindValidate_ParallelExecute(t *testing.T) {
	const (
		workersNum     = 5
		expectedSuffix = "expected one of [Map, Array]"
	)

	wg := sync.WaitGroup{}
	wg.Add(workersNum)

	validate := node.NewKindValidate(node.Map, node.Array)
	nd := node.EmptyNode{}

	for range workersNum {
		go func() {
			defer wg.Done()

			err := validate.Validate(nd)
			require.True(t, strings.HasSuffix(err.Error(), expectedSuffix))
		}()
	}

	wg.Wait()
}
