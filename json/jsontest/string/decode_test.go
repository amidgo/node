package string_test

import (
	_ "embed"
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

var (
	//go:embed testdata/valid/string_array.json
	stringArrayJSON string
	//go:embed testdata/valid/string_single.json
	stringSingleJSON string
	//go:embed testdata/not_valid/string_array_invalid.txt
	stringArrayUnexpectedByteJSON string
	//go:embed testdata/not_valid/string_duo.txt
	duoContentableNodeIsNotInitializedString string
	//go:embed testdata/not_valid/string_invalid_quote.txt
	invalidQuoteString string
	//go:embed testdata/not_valid/non_string_key.txt
	nonStringKey string
)

//nolint:dupl //it is not duplicate
func Test_String_Decode_Success(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName: "string strong object case",
			Data:     stringArrayJSON,
			ExpectedNode: node.MakeArrayNodeWithContent(
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
		&jsontest.DecodeTestCase{
			CaseName:     "single string",
			Data:         stringSingleJSON,
			ExpectedNode: node.MakeStringNode("Hello World!"),
		},
	)
}

func Test_String_Decode_Failure(t *testing.T) {
	tester.RunNamedTesters(
		t,
		&jsontest.DecodeTestCase{
			CaseName:    "not valid array",
			Data:        stringArrayUnexpectedByteJSON,
			ExpectedErr: json.ErrUnexpectedByte,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "not valid duo",
			Data:        duoContentableNodeIsNotInitializedString,
			ExpectedErr: json.ErrUnexpectedByte,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "invalid quote",
			Data:        invalidQuoteString,
			ExpectedErr: json.ErrUnexpectedByte,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "non string key in map",
			Data:        nonStringKey,
			ExpectedErr: json.ErrMissingProperty,
		},
	)
}
