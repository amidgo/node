package node

type Decoder interface {
	Decode(data []byte) (Node, error)
}
