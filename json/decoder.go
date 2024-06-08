package json

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/amidgo/node"
)

type Decoder struct{}

func (d *Decoder) Decode(data []byte) (node.Node, error) {
	return d.DecodeFrom(bytes.NewBuffer(data))
}

func (d *Decoder) DecodeFrom(r io.Reader) (node.Node, error) {
	var reader Reader
	switch r := r.(type) {
	case *bytes.Buffer:
		reader = r
	case *strings.Reader:
		reader = r
	default:
		reader = bufio.NewReader(r)
	}

	scan := newScan(reader)

	return scan.Node()
}

func Decode(data []byte) (node.Node, error) {
	dec := new(Decoder)

	return dec.Decode(data)
}
