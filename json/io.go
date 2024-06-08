package json

import "io"

type ByteRead struct {
	reader io.Reader
}

func ByteReader(r io.Reader) *ByteRead {
	return &ByteRead{reader: r}
}

//nolint:wrapcheck // no need wrap error in decorator
func (b *ByteRead) ReadByte() (byte, error) {
	bt := make([]byte, 1)

	_, err := b.reader.Read(bt)

	return bt[0], err
}

type SkipByteRead struct {
	byteReader io.ByteReader
	skipBytes  []byte
}

func SkipByteReader(byteReader io.ByteReader, skipBytes []byte) *SkipByteRead {
	return &SkipByteRead{
		byteReader: byteReader,
		skipBytes:  skipBytes,
	}
}

//nolint:wrapcheck // no need wrap error in decorator
func (b *SkipByteRead) ReadByte() (byte, error) {
	bt, err := b.byteReader.ReadByte()
	if err != nil {
		return bt, err
	}

	for _, skipByte := range b.skipBytes {
		if bt == skipByte {
			return b.ReadByte()
		}
	}

	return bt, nil
}
