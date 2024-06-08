package json

import (
	"bytes"
	"errors"
	"io"

	"github.com/amidgo/node"
)

var (
	ErrScanMapKey   = errors.New("scan map key")
	ErrScanMapValue = errors.New("scan map value")
)

type mapScan struct {
	reader Reader
}

func (s *mapScan) Node() (node.Node, error) {
	mapNode := node.MakeMapNode()

	byteRead := SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'})

Loop:
	for {
		nextByte, err := byteRead.ReadByte()
		switch {
		case errors.Is(err, io.EOF):
			return nil, ErrMapNotClosed
		case err != nil:
			return nil, errors.Join(ErrRead, err)
		}

		switch nextByte {
		case '}':
			break Loop
		case '"':
		default:
			return nil, NewErrUnexpectedByte(nextByte)
		}

		keyNode, err := s.keyNode()
		if err != nil {
			return nil, err
		}

		valueNode, hasNext, err := s.valueNode()
		if err != nil {
			return nil, err
		}

		mapNode = node.MapAppend(mapNode, keyNode, valueNode)

		if !hasNext {
			break Loop
		}
	}

	return mapNode, nil
}

func (s *mapScan) keyNode() (node.Node, error) {
	mapScan := mapKeyScan{
		scan: stringScan{
			byteReader: s.reader,
		},
		byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
	}

	nd, err := mapScan.Node()
	if err != nil {
		return nil, err
	}

	return nd, nil
}

//nolint:wrapcheck // no need wrap err in contentNodeScan
func (s *mapScan) valueNode() (nd node.Node, hasNext bool, err error) {
	byteRead := SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'})

	nextByte, err := byteRead.ReadByte()
	if err != nil {
		return nil, false, errors.Join(ErrRead)
	}

	valueScan, err := s.valueScan(nextByte)
	if err != nil {
		return nil, false, err
	}

	nd, hasNext, err = valueScan.Node()
	if err != nil {
		return nil, false, err
	}

	return nd, hasNext, err
}

func (s *mapScan) valueScan(nextByte byte) (contentValueScan, error) {
	switch nextByte {
	case '"':
		return &mapValueScan{
			scan: &stringScan{
				byteReader: s.reader,
			},
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
		}, nil
	case '{':
		return &mapValueScan{
			scan: &mapScan{
				reader: s.reader,
			},
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
		}, nil
	case '[':
		return &mapValueScan{
			scan: &arrayScan{
				reader: s.reader,
			},
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
		}, nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		return &numberMapScan{
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
		return &mapValueScan{
			scan: &nullScan{
				reader: s.reader,
			},
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
		}, nil
	case 't':
		return &mapValueScan{
			scan: &trueScan{
				reader: s.reader,
			},
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
		}, nil
	case 'f':
		return &mapValueScan{
			scan: &falseScan{
				reader: s.reader,
			},
			byteReader: SkipByteReader(s.reader, []byte{' ', '\t', '\r', '\n'}),
		}, nil
	default:
		return nil, NewErrUnexpectedByte(nextByte)
	}
}

type mapKeyScan struct {
	scan       stringScan
	byteReader io.ByteReader
}

func (s *mapKeyScan) Node() (node.Node, error) {
	nd, err := s.scan.Node()
	if err != nil {
		return nil, errors.Join(ErrScanMapKey, err)
	}

	nextByte, err := s.byteReader.ReadByte()
	if err != nil {
		return nil, errors.Join(ErrRead, err)
	}

	if nextByte != ':' {
		return nil, NewErrUnexpectedByte(nextByte)
	}

	return nd, nil
}

type mapValueScan struct {
	byteReader io.ByteReader
	scan       scan
}

func (s *mapValueScan) Node() (nd node.Node, hasNext bool, err error) {
	nd, err = s.scan.Node()
	if err != nil {
		return nil, false, errors.Join(ErrScanMapValue, err)
	}

	bt, err := s.byteReader.ReadByte()

	switch {
	case errors.Is(err, io.EOF):
		return nil, false, ErrMapNotClosed
	case err != nil:
		return nil, false, errors.Join(ErrRead, err)
	}

	switch bt {
	case '}':
		return nd, false, nil
	case ',':
		return nd, true, nil
	default:
		return nil, false, NewErrUnexpectedByte(bt)
	}
}

type numberMapScan struct {
	byteReader io.ByteReader
}

func (s *numberMapScan) Node() (nd node.Node, hasNext bool, err error) {
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
		case '}':
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
