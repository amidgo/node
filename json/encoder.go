package json

import (
	"bytes"

	"github.com/amidgo/node"
)

type Encoder struct{}

func (e *Encoder) Encode(nd node.Node) ([]byte, error) {
	output := bytes.Buffer{}
	writer := makeNodeWriterWithOutput(&output)

	err := writer.writeNode(nd)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

func Encode(nd node.Node) ([]byte, error) {
	enc := new(Encoder)

	return enc.Encode(nd)
}
