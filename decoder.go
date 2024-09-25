package node

import (
	"io"
	"log"
)

type Decoder interface {
	Decode(data []byte) (Node, error)
}

type DecoderFrom interface {
	DecodeFrom(src io.Reader) (Node, error)
}

func MustDecode(decoder Decoder, data []byte) Node {
	nd, err := decoder.Decode(data)
	if err != nil {
		log.Panicf("decoder.Decode failed, %s", err)
	}

	return nd
}

func MustDecodeFrom(decoder DecoderFrom, src io.Reader) Node {
	nd, err := decoder.DecodeFrom(src)
	if err != nil {
		log.Panicf("decoder.DecodeFrom failed, %s", err)
	}

	return nd
}
