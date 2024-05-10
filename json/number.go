package json

import "strconv"

//nolint:cyclop //switch case cyclomatic complexity
func numberData(data []byte) ([]byte, bool) {
	var hasE, afterE, hasDot bool

	end := 1

Loop:
	for _, b := range data[end:] {
		switch {
		case b >= '0' && b <= '9':
			afterE = false
		case b == '.' && !hasDot:
			hasDot = true
		case (b == 'e' || b == 'E') && !hasE:
			hasE = true
			hasDot = true
			afterE = true
		case (b == '+' || b == '-') && afterE:
			afterE = false
		default:
			if !isTokenEnd(b) {
				return nil, false
			}

			break Loop
		}

		end++
	}

	return data[:end], true
}

func isTokenEnd(b byte) bool {
	return b == ' ' ||
		b == '\t' ||
		b == '\r' ||
		b == '\n' ||
		b == '[' ||
		b == ']' ||
		b == '{' ||
		b == '}' ||
		b == ',' ||
		b == ':'
}

func tryScanInteger(data string) (int64, bool) {
	value, err := strconv.ParseInt(data, 10, 64)

	return value, err == nil
}

func tryScanFloat(data string) (float64, bool) {
	value, err := strconv.ParseFloat(data, 64)

	return value, err == nil
}
