package node

type Encoder interface {
	Encode(nd Node) ([]byte, error)
}
