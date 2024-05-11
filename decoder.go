package node

import "io"

type Decoder interface {
	Decode(data []byte) (Node, error)
}

type DecoderFrom interface {
	DecodeFrom(src io.Reader) (Node, error)
}
