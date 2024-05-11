package string_test

import (
	"testing"

	_ "embed"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

var (
	//go:embed testdata/valid/string_array_flat.json
	flatStringArrayJSON string
	//go:embed testdata/valid/string_array_indent.json
	indentStringArrayJSON string
)

//nolint:dupl //it is not duplicate
func Test_String_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "string strong object case",
			ExpectedData: flatStringArrayJSON,
			Node: node.MakeArrayNodeWithContent(
				node.MakeMapNodeWithContent(
					node.MakeStringNode("name"),
					node.MakeStringNode("Dima"),
					node.MakeStringNode("birthdate"),
					node.MakeStringNode("2004-11-25"),
				),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("name"),
					node.MakeStringNode("Vika"),
					node.MakeStringNode("birthdate"),
					node.MakeStringNode("2009-05-28"),
				),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("name"),
					node.MakeStringNode("иван"),
					node.MakeStringNode("birthdate"),
					node.MakeStringNode("иван"),
				),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("name"),
					node.MakeStringNode("андрей"),
					node.MakeStringNode("birthdate"),
					node.MakeStringNode("1973-04-25"),
				),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("user"),
					node.MakeMapNodeWithContent(
						node.MakeStringNode("name"),
						node.MakeStringNode("андрей"),
						node.MakeStringNode("birthdate"),
						node.MakeStringNode("1973-04-25"),
					),
				),
				node.MakeStringNode("hello"),
				node.MakeStringNode("aboba"),
				node.MakeArrayNodeWithContent(
					node.MakeStringNode("hello"),
					node.MakeStringNode("aboba"),
				),
			),
		},
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
			Node:         node.MakeStringNode(`"Hello World"`),
			ExpectedData: `""Hello World""`,
		},
	)

}

func Test_EncodeIndent(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "string strong object case",
			ExpectedData: indentStringArrayJSON,
			Indent:       4,
			Node: node.MakeArrayNodeWithContent(
				node.MakeMapNodeWithContent(
					node.MakeStringNode("name"),
					node.MakeStringNode("Dima"),
					node.MakeStringNode("birthdate"),
					node.MakeStringNode("2004-11-25"),
				),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("name"),
					node.MakeStringNode("Vika"),
					node.MakeStringNode("birthdate"),
					node.MakeStringNode("2009-05-28"),
				),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("name"),
					node.MakeStringNode("иван"),
					node.MakeStringNode("birthdate"),
					node.MakeStringNode("иван"),
				),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("name"),
					node.MakeStringNode("андрей"),
					node.MakeStringNode("birthdate"),
					node.MakeStringNode("1973-04-25"),
				),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("user"),
					node.MakeMapNodeWithContent(
						node.MakeStringNode("name"),
						node.MakeStringNode("андрей"),
						node.MakeStringNode("birthdate"),
						node.MakeStringNode("1973-04-25"),
					),
				),
				node.MakeStringNode("hello"),
				node.MakeStringNode("aboba"),
				node.MakeArrayNodeWithContent(
					node.MakeStringNode("hello"),
					node.MakeStringNode("aboba"),
				),
			),
		},
	)
}
