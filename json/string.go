package json

import (
	"errors"
	"io"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/amidgo/node"
)

var (
	ErrIncorrectEscapedSymbol   = errors.New("incorrect escape symbol \\ at the end of token")
	ErrIncorrectEscapedBytes    = errors.New("incorrectly escaped bytes")
	ErrIncorrectEscapedSequence = errors.New("incorrectly escaped \\uXXXX sequence")
	ErrStringNotValid           = errors.New("string not valid")
)

type stringScan struct {
	byteReader io.ByteReader
}

func (s *stringScan) Node() (node.Node, error) {
	data, err := s.stringData()
	if err != nil {
		return nil, err
	}

	sv, err := stringValue(data)
	if err != nil {
		return nil, err
	}

	return node.MakeStringNode(sv), nil
}

func (s *stringScan) stringData() ([]byte, error) {
	data := make([]byte, 0)

	var skip bool

	for {
		b, err := s.byteReader.ReadByte()
		if err != nil {
			return nil, errors.Join(ErrRead, err)
		}

		data = append(data, b)

		if skip {
			skip = false

			continue
		}

		switch b {
		case '\\':
			skip = true
		case '"':
			return data, nil
		}
	}
}

func stringValue(data []byte) (string, error) {
	res, ok := unquoteBytes(data)
	if !ok {
		return "", ErrStringNotValid
	}

	return string(res), nil
}

//nolint:funlen,gocognit // function from go sources
func unquoteBytes(s []byte) (t []byte, ok bool) {
	if len(s) < 1 || s[len(s)-1] != '"' {
		return t, ok
	}

	s = s[:len(s)-1]

	// Check for unusual characters. If there are none,
	// then no unquoting is needed, so return a slice of the
	// original bytes.
	r := 0
	for r < len(s) {
		c := s[r]
		if c == '\\' || c == '"' || c < ' ' {
			break
		}

		if c < utf8.RuneSelf {
			r++

			continue
		}

		rr, size := utf8.DecodeRune(s[r:])
		if rr == utf8.RuneError && size == 1 {
			break
		}

		r += size
	}

	if r == len(s) {
		return s, true
	}

	b := make([]byte, len(s)+2*utf8.UTFMax)
	w := copy(b, s[0:r])

	for r < len(s) {
		// Out of room? Can only happen if s is full of
		// malformed UTF-8 and we're replacing each
		// byte with RuneError.
		if w >= len(b)-2*utf8.UTFMax {
			nb := make([]byte, (len(b)+utf8.UTFMax)*2)
			copy(nb, b[0:w])
			b = nb
		}

		switch c := s[r]; {
		case c == '\\':
			r++
			if r >= len(s) {
				return t, ok
			}

			switch s[r] {
			default:
				return t, ok
			case '"', '\\', '/', '\'':
				b[w] = s[r]
				r++
				w++
			case 'b':
				b[w] = '\b'
				r++
				w++
			case 'f':
				b[w] = '\f'
				r++
				w++
			case 'n':
				b[w] = '\n'
				r++
				w++
			case 'r':
				b[w] = '\r'
				r++
				w++
			case 't':
				b[w] = '\t'
				r++
				w++
			case 'u':
				r--
				rr := getu4(s[r:])

				if rr < 0 {
					return t, ok
				}

				r += 6

				if utf16.IsSurrogate(rr) {
					rr1 := getu4(s[r:])
					if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
						// A valid pair; consume.
						r += 6
						w += utf8.EncodeRune(b[w:], dec)

						break
					}
					// Invalid surrogate; fall back to replacement rune.
					rr = unicode.ReplacementChar
				}

				w += utf8.EncodeRune(b[w:], rr)
			}

		// Quote, control characters are invalid.
		case c == '"', c < ' ':
			return t, ok

		// ASCII
		case c < utf8.RuneSelf:
			b[w] = c
			r++
			w++

		// Coerce to well-formed UTF-8.
		default:
			rr, size := utf8.DecodeRune(s[r:])
			r += size
			w += utf8.EncodeRune(b[w:], rr)
		}
	}

	return b[0:w], true
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.
func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}

	var r rune

	for _, c := range s[2:6] {
		switch {
		case '0' <= c && c <= '9':
			c -= '0'
		case 'a' <= c && c <= 'f':
			c = c - 'a' + 10
		case 'A' <= c && c <= 'F':
			c = c - 'A' + 10
		default:
			return -1
		}

		r = r*16 + rune(c)
	}

	return r
}
