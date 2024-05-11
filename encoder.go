package node

import (
	"io"
)

type Encoder interface {
	Encode(nd Node) ([]byte, error)
}

type EncoderTo interface {
	EncodeTo(w io.Writer, node Node) error
}
