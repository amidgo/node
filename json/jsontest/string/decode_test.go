package string_test

import (
	_ "embed"
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Decode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName:     "single string",
			Data:         `"Hello World!"`,
			ExpectedNode: node.MakeStringNode("Hello World!"),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "single quoted string",
			Data:         `"'Hello World'"`,
			ExpectedNode: node.MakeStringNode(`'Hello World'`),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "double quoted string",
			Data:         `"\"Hello World!\""`,
			ExpectedNode: node.MakeStringNode(`"Hello World!"`),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "unicode string",
			Data:         `"\u0048\u0065\u006C\u006C\u006F\u0020\u0057\u006F\u0072\u006C\u0064\u0021"`,
			ExpectedNode: node.MakeStringNode("Hello World!"),
		},
		&jsontest.DecodeTestCase{
			CaseName:    "invalid string",
			Data:        `Hello World!"`,
			ExpectedErr: json.NewErrUnexpectedByte('H'),
		},
		&jsontest.DecodeTestCase{
			CaseName:    "invalid unicode string",
			Data:        `"\u08\u0065\u006C\u006C\u006F\u0020\u0057\u006F\u0072\u006C\u0064\u0021"`,
			ExpectedErr: json.ErrStringNotValid,
		},
	)
}
