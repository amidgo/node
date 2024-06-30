package node

type Validate interface {
	Validate(nd Node) error
}
