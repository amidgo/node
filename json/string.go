package json

import (
	"bytes"
	"errors"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

var (
	ErrIncorrectEscapedSymbol   = errors.New("incorrect escape symbol \\ at the end of token")
	ErrIncorrectEscapedBytes    = errors.New("incorrectly escaped bytes")
	ErrIncorrectEscapedSequence = errors.New("incorrectly escaped \\uXXXX sequence")
)

func findStringLen(data []byte) (isValid bool, length int) {
	for {
		idx := bytes.IndexByte(data, '"')
		if idx == -1 {
			return false, len(data)
		}

		if idx == 0 || (idx > 0 && data[idx-1] != '\\') {
			return true, length + idx
		}

		// count \\\\\\\ sequences. even number of slashes means quote is not really escaped
		cnt := 1
		for idx-cnt-1 >= 0 && data[idx-cnt-1] == '\\' {
			cnt++
		}

		if cnt%2 == 0 {
			return true, length + idx
		}

		length += idx + 1
		data = data[idx+1:]
	}
}

func stringValue(data []byte) (string, error) {
	res, err := unescapeStringToken(data)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func unescapeStringToken(data []byte) ([]byte, error) {
	res := data

	var unescapedData []byte

	for {
		i := bytes.IndexByte(data, '\\')
		if i == -1 {
			break
		}

		escapedRune, escapedBytes, err := decodeEscape(data[i:])
		if err != nil {
			return nil, err
		}

		if unescapedData == nil {
			unescapedData = make([]byte, 0, len(res))
		}

		var d [4]byte
		s := utf8.EncodeRune(d[:], escapedRune)

		unescapedData = append(unescapedData, data[:i]...)
		unescapedData = append(unescapedData, d[:s]...)

		data = data[i+escapedBytes:]
	}

	if unescapedData != nil {
		//nolint:gocritic // append to result with if statement
		res = append(unescapedData, data...)
	}

	return res, nil
}

//nolint:gomnd // 2 bytes is default bytes processed size
func decodeEscape(data []byte) (decoded rune, bytesProcessed int, err error) {
	if len(data) < 2 {
		return 0, 0, ErrIncorrectEscapedSymbol
	}

	c := data[1]
	switch c {
	case '"', '/', '\\':
		return rune(c), 2, nil
	case 'b':
		return '\b', 2, nil
	case 'f':
		return '\f', 2, nil
	case 'n':
		return '\n', 2, nil
	case 'r':
		return '\r', 2, nil
	case 't':
		return '\t', 2, nil
	case 'u':
		return decodeUnicodeEscapeData(data)
	}

	return 0, 0, ErrIncorrectEscapedBytes
}

func decodeUnicodeEscapeData(data []byte) (decoded rune, bytesProcessed int, err error) {
	rr := getu4(data)
	if rr < 0 {
		return 0, 0, ErrIncorrectEscapedSequence
	}

	read := 6
	if utf16.IsSurrogate(rr) {
		rr1 := getu4(data[read:])
		if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
			read += 6
			rr = dec
		} else {
			rr = unicode.ReplacementChar
		}
	}

	return rr, read, nil
}

//nolint:gomnd //calculated operations with bytes
func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}

	var val rune

	for i := 2; i < len(s) && i < 6; i++ {
		var v byte

		c := s[i]

		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v = c - '0'
		case 'a', 'b', 'c', 'd', 'e', 'f':
			v = c - 'a' + 10
		case 'A', 'B', 'C', 'D', 'E', 'F':
			v = c - 'A' + 10
		default:
			return -1
		}

		val <<= 4
		val |= rune(v)
	}

	return val
}
