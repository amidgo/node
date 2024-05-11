package json

import (
	"errors"
	"fmt"
	"io"

	"github.com/amidgo/node"
)

type Decoder struct{}

func (d *Decoder) Decode(data []byte) (node.Node, error) {
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

func (d *Decoder) DecodeFrom(r io.Reader) (node.Node, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read err: %w", err)
	}

	return d.Decode(data)
}

func Decode(data []byte) (node.Node, error) {
	dec := new(Decoder)

	return dec.Decode(data)
}
