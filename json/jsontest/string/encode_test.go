package string_test

import (
	"testing"

	_ "embed"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_String_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "single string",
			Node:         node.MakeStringNode("Hello World!"),
			ExpectedData: `"Hello World!"`,
		},
		&jsontest.EncodeTestCase{
			CaseName:     "single quoted string",
			Node:         node.MakeStringNode(`'Hello World'`),
			ExpectedData: `"'Hello World'"`,
		},
		&jsontest.EncodeTestCase{
			CaseName:     "double quoted string",
			Node:         node.MakeStringNode(`"Hello World!"`),
			ExpectedData: "\"\"Hello World!\"\"",
		},
		&jsontest.EncodeTestCase{
			CaseName:     "unicode string",
			Node:         node.MakeStringNode("\u0048\u0065\u006C\u006C\u006F\u0020\u0057\u006F\u0072\u006C\u0064\u0021"),
			ExpectedData: `"Hello World!"`,
		},
	)
}
