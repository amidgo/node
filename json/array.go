package json

import (
	"bytes"
	"errors"
	"io"

	"github.com/amidgo/node"
)

var (
	ErrArrayNotClosed  = errors.New("array not closed")
	ErrScanNodeInArray = errors.New("scan node in array")
)

type arrayScan struct {
	reader Reader
}

func (s *arrayScan) Node() (node.Node, error) {
	arrayNode := node.MakeArrayNode()

	byteRead := SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'})

Loop:
	for {
		nextByte, err := byteRead.ReadByte()
		switch {
		case errors.Is(err, io.EOF):
			return nil, ErrArrayNotClosed
		case err != nil:
			return nil, errors.Join(ErrRead, err)
		}

		if nextByte == ']' {
			break Loop
		}

		scan, err := s.nodeScan(nextByte)
		if err != nil {
			return nil, err
		}

		nd, hasNext, err := scan.Node()
		if err != nil {
			return nil, errors.Join(ErrScanNodeInArray, err)
		}

		arrayNode = node.ArrayAppend(arrayNode, nd)

		if !hasNext {
			break Loop
		}
	}

	return arrayNode, nil
}

func (s *arrayScan) nodeScan(nextByte byte) (contentValueScan, error) {
	switch nextByte {
	case '"':
		return &arrayValueScan{
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
			scan: &stringScan{
				byteReader: s.reader,
			},
		}, nil
	case '{':
		return &arrayValueScan{
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
			scan: &mapScan{
				reader: s.reader,
			},
		}, nil
	case '[':
		return &arrayValueScan{
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
			scan: &arrayScan{
				reader: s.reader,
			},
		}, nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		return &numberArrayScan{
			byteReader: SkipByteReader(
				ByteReader(
					io.MultiReader(
						bytes.NewBuffer([]byte{nextByte}),
						s.reader,
					),
				),
				[]byte{' ', '\t', '\r', '\n'},
			),
		}, nil
	case 'n':
		return &arrayValueScan{
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
			scan: &nullScan{
				reader: s.reader,
			},
		}, nil
	case 't':
		return &arrayValueScan{
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
			scan: &trueScan{
				reader: s.reader,
			},
		}, nil
	case 'f':
		return &arrayValueScan{
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
			scan: &falseScan{
				reader: s.reader,
			},
		}, nil
	default:
		return nil, NewErrUnexpectedByte(nextByte)
	}
}

type arrayValueScan struct {
	byteReader io.ByteReader
	scan       scan
}

//nolint:wrapcheck // no wrap err recursively
func (s *arrayValueScan) Node() (nd node.Node, hasNext bool, err error) {
	nd, err = s.scan.Node()
	if err != nil {
		return nil, false, err
	}

	bt, err := s.byteReader.ReadByte()

	switch {
	case errors.Is(err, io.EOF):
		return nil, false, ErrArrayNotClosed
	case err != nil:
		return nil, false, errors.Join(ErrRead, err)
	}

	switch bt {
	case ']':
		return nd, false, nil
	case ',':
		return nd, true, nil
	default:
		return nil, false, NewErrUnexpectedByte(bt)
	}
}

type numberArrayScan struct {
	byteReader io.ByteReader
}

func (s *numberArrayScan) Node() (nd node.Node, hasNext bool, err error) {
	buf := &bytes.Buffer{}

Loop:
	for {
		bt, err := s.byteReader.ReadByte()
		switch {
		case errors.Is(err, io.EOF):
			break Loop
		case err != nil:
			return nil, false, errors.Join(ErrRead, err)
		}

		switch bt {
		case ']':
			hasNext = false

			break Loop
		case ',':
			hasNext = true

			break Loop
		default:
			buf.WriteByte(bt)
		}
	}

	scan := numberScan{
		byteReader: buf,
	}

	nd, err = scan.Node()
	if err != nil {
		return nil, false, err
	}

	return nd, hasNext, nil
}
