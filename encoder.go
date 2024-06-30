package node

import (
	"io"
	"log"
)

type Encoder interface {
	Encode(nd Node) ([]byte, error)
}

type EncoderTo interface {
	EncodeTo(w io.Writer, node Node) error
}

func MustEncode(encoder Encoder, nd Node) []byte {
	data, err := encoder.Encode(nd)
	if err != nil {
		log.Panicf("encoder.Encode failed, %s", err)
	}

	return data
}

func MustEncodeTo(encoder EncoderTo, dst io.Writer, nd Node) {
	err := encoder.EncodeTo(dst, nd)
	if err != nil {
		log.Panicf("encoder.EncodeTo failed, %s", err)
	}
}
