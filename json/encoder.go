package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/amidgo/node"
)

type Encoder struct {
	Indent int
}

func (e *Encoder) Encode(nd node.Node) ([]byte, error) {
	output := bytes.Buffer{}

	err := e.EncodeTo(&output, nd)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

func (e *Encoder) EncodeTo(w io.Writer, nd node.Node) error {
	if e.Indent > 0 {
		return e.encodeIndent(w, nd)
	}

	writer := makeNodeWriterWithOutput(w)

	err := writer.writeNode(nd)
	if err != nil {
		return err
	}

	return nil
}

func (e *Encoder) encodeIndent(output io.Writer, nd node.Node) error {
	notIndentJson := &bytes.Buffer{}

	writer := makeNodeWriterWithOutput(notIndentJson)

	err := writer.writeNode(nd)
	if err != nil {
		return err
	}

	indentedJson := &bytes.Buffer{}

	bufOutput, ok := output.(*bytes.Buffer)
	if ok {
		indentedJson = bufOutput
	}

	err = json.Indent(
		indentedJson,
		notIndentJson.Bytes(), "",
		strings.Repeat(" ", e.Indent),
	)
	if err != nil {
		return fmt.Errorf("indenting json, %w", err)
	}

	if !ok {
		_, err := indentedJson.WriteTo(output)
		if err != nil {
			return fmt.Errorf("write indented to output, %w", err)
		}
	}

	return nil
}

func Encode(nd node.Node) ([]byte, error) {
	enc := new(Encoder)

	return enc.Encode(nd)
}
