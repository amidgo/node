package json

import (
	"bytes"
	"errors"
	"io"
	"strconv"

	"github.com/amidgo/node"
)

var ErrNumberNotValid = errors.New("number not valid")

type numberScan struct {
	byteReader io.ByteReader
}

func (s *numberScan) Node() (node.Node, error) {
	data, err := s.numberData()
	if err != nil {
		return nil, errors.Join(ErrNumberNotValid, err)
	}

	i, ok := tryScanInteger(data)
	if ok {
		return node.MakeIntegerNode(i), nil
	}

	f, ok := tryScanFloat(data)
	if ok {
		return node.MakeFloatNode(f), nil
	}

	return nil, ErrNumberNotValid
}

func (s *numberScan) numberData() (string, error) {
	data := bytes.Buffer{}

Loop:
	for {
		b, err := s.byteReader.ReadByte()
		switch {
		case errors.Is(err, io.EOF):
			break Loop
		case err != nil:
			return "", errors.Join(ErrRead, err)
		}

		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
			'.', 'e', 'E', '+', '-':
			data.WriteByte(b)
		default:
			return "", NewErrUnexpectedByte(b)
		}
	}

	return data.String(), nil
}

func tryScanInteger(data string) (int64, bool) {
	value, err := strconv.ParseInt(data, 10, 64)

	return value, err == nil
}

func tryScanFloat(data string) (float64, bool) {
	value, err := strconv.ParseFloat(data, 64)

	return value, err == nil
}
