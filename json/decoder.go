package json

import (
	"errors"
	"io"

	"github.com/amidgo/node"
)

type Decoder struct{}

func (p *Decoder) Decode(data []byte) (node.Node, error) {
	scanner := newScanner(data)
	for scanner.HasNext() {
		err := scanner.Scan()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}
	}

	return scanner.Node(), nil
}

func Decode(data []byte) (node.Node, error) {
	dec := new(Decoder)

	return dec.Decode(data)
}
