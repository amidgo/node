package json

import (
	"bytes"
	"errors"
	"io"
	"slices"

	"github.com/amidgo/node"
)

var (
	ErrRead          = errors.New("read data")
	ErrMapNotClosed  = errors.New("map not closed")
	ErrNullNotValid  = errors.New("null not valid")
	ErrTrueNotValid  = errors.New("true not valid")
	ErrFalseNotValid = errors.New("false not valid")
)

type UnexpectedByteError struct {
	value byte
}

func NewErrUnexpectedByte(value byte) UnexpectedByteError {
	return UnexpectedByteError{value: value}
}

func (e UnexpectedByteError) Error() string {
	return "receive unexpected byte " + string([]byte{e.value})
}

func (e UnexpectedByteError) Is(err error) bool {
	unexp, ok := err.(UnexpectedByteError)
	if !ok {
		return false
	}

	return unexp.value == e.value
}

type Reader interface {
	io.Reader
	io.ByteReader
}

type scan interface {
	Node() (node.Node, error)
}

type contentValueScan interface {
	Node() (node node.Node, hasNext bool, err error)
}

type rootScan struct {
	reader Reader
}

func newScan(reader Reader) rootScan {
	return rootScan{reader: reader}
}

//nolint:wrapcheck // no need to wrap Scan err recursively
func (s *rootScan) Node() (node.Node, error) {
	nodeScan, err := s.nodeScan()
	if err != nil {
		return nil, err
	}

	nd, err := nodeScan.Node()
	if err != nil {
		return nil, err
	}

	return nd, nil
}

func (s *rootScan) nodeScan() (scan, error) {
	read := SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'})

	nextByte, err := read.ReadByte()
	if err != nil {
		return nil, errors.Join(ErrRead, err)
	}

	switch nextByte {
	case '"':
		return &stringScan{byteReader: s.reader}, nil
	case '{':
		return &mapScan{reader: s.reader}, nil
	case '[':
		return &arrayScan{reader: s.reader}, nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		return &numberScan{
			byteReader: ByteReader(
				io.MultiReader(
					bytes.NewReader([]byte{nextByte}),
					s.reader,
				),
			),
		}, nil
	case 'n':
		return &nullScan{reader: s.reader}, nil
	case 't':
		return &trueScan{reader: s.reader}, nil
	case 'f':
		return &falseScan{reader: s.reader}, nil
	default:
		return nil, NewErrUnexpectedByte(nextByte)
	}
}

type nullScan struct {
	reader io.Reader
}

func (s *nullScan) Node() (node.Node, error) {
	data := make([]byte, 3)

	_, err := s.reader.Read(data)
	if err != nil {
		return nil, errors.Join(ErrNullNotValid, ErrRead, err)
	}

	if slices.Equal(data, []byte{'u', 'l', 'l'}) {
		return node.EmptyNode{}, nil
	}

	return nil, ErrNullNotValid
}

type trueScan struct {
	reader io.Reader
}

func (s *trueScan) Node() (node.Node, error) {
	data := make([]byte, 3)

	_, err := s.reader.Read(data)
	if err != nil {
		return nil, errors.Join(ErrTrueNotValid, ErrRead, err)
	}

	if slices.Equal(data, []byte{'r', 'u', 'e'}) {
		return node.MakeBoolNode(true), nil
	}

	return nil, ErrTrueNotValid
}

type falseScan struct {
	reader io.Reader
}

func (s *falseScan) Node() (node.Node, error) {
	data := make([]byte, 4)

	_, err := s.reader.Read(data)
	if err != nil {
		return nil, errors.Join(ErrFalseNotValid, ErrRead, err)
	}

	if slices.Equal(data, []byte{'a', 'l', 's', 'e'}) {
		return node.MakeBoolNode(false), nil
	}

	return nil, ErrFalseNotValid
}
