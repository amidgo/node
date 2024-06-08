package number_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Decode_Number(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName:     "1234 integer",
			Data:         `1234`,
			ExpectedNode: node.MakeIntegerNode(1234),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "negative integer",
			Data:         "-2138315144744848",
			ExpectedNode: node.MakeIntegerNode(-2138315144744848),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "1234.5 float",
			Data:         "1234.5",
			ExpectedNode: node.MakeFloatNode(1234.5),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "-1234.5 float",
			Data:         "-1234.5",
			ExpectedNode: node.MakeFloatNode(-1234.5),
		},
		&jsontest.DecodeTestCase{
			CaseName:    "invalid integer",
			Data:        "1.ru",
			ExpectedErr: json.NewErrUnexpectedByte('r'),
		},
	)
}
