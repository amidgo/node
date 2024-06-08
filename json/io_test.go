package json_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/amidgo/node/json"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type SkipByteReadTest struct {
	CaseName   string
	SkipBytes  []byte
	ByteReader io.ByteReader

	ExpectedByte byte
	ExpectedErr  error
}

func (s *SkipByteReadTest) Name() string {
	return s.CaseName
}

func (s *SkipByteReadTest) Test(t *testing.T) {
	skipByteRead := json.SkipByteReader(s.ByteReader, s.SkipBytes)

	nextByte, err := skipByteRead.ReadByte()
	require.ErrorIs(t, err, s.ExpectedErr)
	assert.Equal(t, s.ExpectedByte, nextByte)
}

func Test_SkipByteReader(t *testing.T) {
	tester.RunNamedTesters(t,
		&SkipByteReadTest{
			CaseName:     "empty",
			SkipBytes:    []byte{},
			ByteReader:   bytes.NewReader(nil),
			ExpectedByte: 0,
			ExpectedErr:  io.EOF,
		},
		&SkipByteReadTest{
			CaseName:     "skip single byte",
			SkipBytes:    []byte{' '},
			ByteReader:   strings.NewReader("  h"),
			ExpectedByte: 'h',
			ExpectedErr:  nil,
		},
		&SkipByteReadTest{
			CaseName:     "skip many bytes",
			SkipBytes:    []byte{' ', 'u'},
			ByteReader:   strings.NewReader(" uuuuuu  h"),
			ExpectedByte: 'h',
			ExpectedErr:  nil,
		},
	)
}
