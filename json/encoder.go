package json

import (
	"bytes"
	"io"

	"github.com/amidgo/node"
)

type Encoder struct{}

func (e *Encoder) Encode(nd node.Node) ([]byte, error) {
	output := bytes.Buffer{}

	err := e.EncodeTo(&output, nd)

	return output.Bytes(), err
}

func (e *Encoder) EncodeTo(w io.Writer, nd node.Node) error {
	writer := makeNodeWriterWithOutput(w)

	err := writer.writeNode(nd)
	if err != nil {
		return err
	}

	return nil
}

func Encode(nd node.Node) ([]byte, error) {
	enc := new(Encoder)

	return enc.Encode(nd)
}
