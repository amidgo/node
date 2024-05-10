package json

type nextByteExpecter interface {
	expectByte(b byte) (skipBytes int, valid bool)
}

type mapNodeByteExpecter struct{}

func (e mapNodeByteExpecter) expectByte(b byte) (int, bool) {
	if b == '"' {
		return 0, true
	}

	return 0, false
}

type arrayNodeByteExpecter struct{}

func (e arrayNodeByteExpecter) expectByte(b byte) (int, bool) {
	switch b {
	case '"', '{', '[', 'n', 'f', 't', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		return 0, true
	default:
		return 0, false
	}
}

type mapKeyNodeByteExpecter struct{}

func (e mapKeyNodeByteExpecter) expectByte(b byte) (int, bool) {
	switch b {
	case ':':
		return 1, true
	default:
		return 0, false
	}
}

type mapValueNodeExpecter struct{}

func (e mapValueNodeExpecter) expectByte(b byte) (int, bool) {
	switch b {
	case ',':
		return 1, true
	default:
		return 0, false
	}
}

type arrayValueNodeExpecter struct{}

func (e arrayValueNodeExpecter) expectByte(b byte) (int, bool) {
	switch b {
	case ',':
		return 1, true
	default:
		return 0, false
	}
}

type nilArrayValueNodeExpecter struct{}

func (e nilArrayValueNodeExpecter) expectByte(_ byte) (int, bool) {
	return 0, false
}
