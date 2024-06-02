package json

import (
	"errors"
	"fmt"
	"io"

	"github.com/amidgo/node"
)

var (
	ErrNoNextNodeToScan                = errors.New("no next node to scan")
	ErrWrongScanState                  = errors.New("wrong scan state")
	ErrObjectWithoutKey                = errors.New("object without key")
	ErrStringNotValid                  = errors.New("string not valid")
	ErrNumberNotValid                  = errors.New("number not valid")
	ErrNullNotValid                    = errors.New("null not valid")
	ErrTrueNotValid                    = errors.New("false not valid")
	ErrFalseNotValid                   = errors.New("false not valid")
	ErrUnexpectedByte                  = errors.New("unexpected byte")
	ErrWrongCloseArrayNode             = errors.New("wrong close array node")
	ErrWrongCloseMapNode               = errors.New("wrong close map node")
	ErrNoneContentableNodeLeft         = errors.New("none contentable node left")
	ErrNoNextNodeScanned               = errors.New("no next node scanned")
	ErrMissingProperty                 = errors.New("missing property for map")
	ErrMissingPropertyValue            = errors.New("missing property value for map")
	ErrContentableNodeIsNotInitialized = errors.New("invalid contentable node is not initialized")
	ErrContentableNodeNotClosed        = errors.New("contentable node is not closed")
)

type scanState int

const (
	empty scanState = iota
	scanMap
	scanArray
	scanString
	scanNumber
	scanTrue
	scanFalse
	scanNull
)

type scanner struct {
	inputData []byte

	node node.Node

	currentContentNode node.Node
	contentableNodes   []node.Node

	nextByteExpecter nextByteExpecter

	scanStart int
	hasNext   bool

	currentLine int
}

func newScanner(inputData []byte) *scanner {
	return &scanner{
		nextByteExpecter: nilArrayValueNodeExpecter{},
		inputData:        inputData,
		contentableNodes: make([]node.Node, 0),
		hasNext:          true,
		currentLine:      1,
	}
}

func (s *scanner) Node() node.Node {
	return s.node
}

func (s *scanner) HasNext() bool {
	return s.hasNext
}

func (s *scanner) Scan() error {
	if !s.HasNext() {
		return ErrNoNextNodeToScan
	}

	state, err := s.fetchScanState()
	if err != nil {
		return err
	}

	err = s.scanNode(state)
	if err != nil {
		return fmt.Errorf("failed scanNode, %w, line %d", err, s.currentLine)
	}

	err = s.next()
	if err != nil {
		return fmt.Errorf("failed next, %w, line %d", err, s.currentLine)
	}

	return nil
}

//nolint:cyclop //switch-case cyclomatic complexity
func (s *scanner) fetchScanState() (scanState, error) {
	for _, b := range s.inputData[s.scanStart:] {
		switch b {
		case ' ', '\t', '\r':
			s.scanStart++
		case '\n':
			s.scanStart++
			s.currentLine++
		case '"':
			s.scanStart++

			return scanString, nil
		case '{':
			s.scanStart++

			return scanMap, nil
		case '[':
			s.scanStart++

			return scanArray, nil
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			return scanNumber, nil
		case 'n':
			s.scanStart++

			return scanNull, nil
		case 't':
			s.scanStart++

			return scanTrue, nil
		case 'f':
			s.scanStart++

			return scanFalse, nil
		default:
			return empty, fmt.Errorf("%w: <%s>", ErrUnexpectedByte, string(b))
		}
	}

	return empty, ErrNoNextNodeScanned
}

func (s *scanner) scanNode(state scanState) error {
	switch state {
	case scanMap:
		return s.scanMap()
	case scanArray:
		return s.scanArray()
	case scanString:
		return s.scanString()
	case scanNumber:
		return s.scanNumber()
	case scanTrue:
		return s.scanTrue()
	case scanFalse:
		return s.scanFalse()
	case scanNull:
		return s.scanNull()
	default:
		return ErrWrongScanState
	}
}

func (s *scanner) scanMap() error {
	s.appendContentableNode(node.MakeMapNode())

	return nil
}

func (s *scanner) scanArray() error {
	s.appendContentableNode(node.MakeArrayNode())

	return nil
}

func (s *scanner) scanString() error {
	data := s.inputData[s.scanStart:]

	valid, length := findStringLen(data)
	if !valid {
		return ErrStringNotValid
	}

	data = data[:length]

	s.scanStart += length + 1

	value, err := stringValue(data)
	if err != nil {
		return err
	}

	err = s.appendValueNode(node.MakeStringNode(value))
	if err != nil {
		return err
	}

	s.setValueNodeExpecter()

	return nil
}

func (s *scanner) scanNumber() error {
	data := s.inputData[s.scanStart:]

	numberData, valid := numberData(data)
	if !valid {
		return ErrNumberNotValid
	}

	numberNode, err := numberNode(string(numberData))
	if err != nil {
		return err
	}

	err = s.appendValueNode(numberNode)
	if err != nil {
		return err
	}

	s.scanStart += len(numberData)

	s.setValueNodeExpecter()

	return nil
}

func numberNode(numberData string) (node.Node, error) {
	i, ok := tryScanInteger(numberData)
	if ok {
		return node.MakeIntegerNode(i), nil
	}

	f, ok := tryScanFloat(numberData)
	if !ok {
		return nil, ErrNumberNotValid
	}

	return node.MakeFloatNode(f), nil
}

func (s *scanner) scanTrue() error {
	const rueLength = 3

	data := s.inputData[s.scanStart:]
	if len(data) < rueLength {
		return ErrTrueNotValid
	}

	if [3]byte(data) != [3]byte{'r', 'u', 'e'} {
		return ErrTrueNotValid
	}

	err := s.appendValueNode(node.MakeBoolNode(true))
	if err != nil {
		return err
	}

	s.scanStart += rueLength

	s.setValueNodeExpecter()

	return nil
}

func (s *scanner) scanFalse() error {
	const alseLength = 4

	data := s.inputData[s.scanStart:]
	if len(data) < alseLength {
		return ErrFalseNotValid
	}

	if [4]byte(data) != [4]byte{'a', 'l', 's', 'e'} {
		return ErrFalseNotValid
	}

	err := s.appendValueNode(node.MakeBoolNode(false))
	if err != nil {
		return err
	}

	s.scanStart += alseLength

	s.setValueNodeExpecter()

	return nil
}

func (s *scanner) scanNull() error {
	const ullLength = 3

	data := s.inputData[s.scanStart:]
	if len(data) < ullLength {
		return ErrNullNotValid
	}

	if [3]byte(data) != [3]byte{'u', 'l', 'l'} {
		return ErrNullNotValid
	}

	err := s.appendValueNode(node.EmptyNode{})
	if err != nil {
		return err
	}

	s.scanStart += ullLength

	s.setValueNodeExpecter()

	return nil
}

func (s *scanner) setValueNodeExpecter() {
	if s.currentContentNode == nil {
		return
	}

	switch s.currentContentNode.Kind() {
	case node.Array:
		s.nextByteExpecter = arrayValueNodeExpecter{}
	case node.Map:
		s.setMapValueNodeExpecter()
	}
}

func (s *scanner) setMapValueNodeExpecter() {
	switch len(s.currentContentNode.Content()) % 2 {
	case 1:
		s.nextByteExpecter = mapKeyNodeByteExpecter{}
	case 0:
		s.nextByteExpecter = mapValueNodeExpecter{}
	}
}

func (s *scanner) next() error {
	for _, b := range s.inputData[s.scanStart:] {
		switch b {
		case ' ', '\t', '\r':
			s.scanStart++
		case '\n':
			s.scanStart++
			s.currentLine++
		case '}':
			s.scanStart++

			err := s.closeMapNode()
			if err != nil {
				return err
			}
		case ']':
			s.scanStart++

			err := s.closeArrayNode()
			if err != nil {
				return err
			}
		default:
			skipBytes, expect := s.nextByteExpecter.expectByte(b)
			if expect {
				s.scanStart += skipBytes

				return nil
			}

			return fmt.Errorf("%w: <%s>", ErrUnexpectedByte, string(b))
		}
	}

	s.hasNext = false

	if !s.allContentableNodesClosed() {
		return ErrContentableNodeNotClosed
	}

	return io.EOF
}

func (s *scanner) appendContentableNode(nd node.Node) {
	if s.node == nil {
		s.node = nd
	} else if s.currentContentNode != nil {
		s.appendNode(nd)
	}

	s.currentContentNode = nd
	s.contentableNodes = append(s.contentableNodes, nd)

	s.setCurrentContentNodeByteExpecter()
}

func (s *scanner) setCurrentContentNodeByteExpecter() {
	if s.currentContentNode == nil {
		return
	}

	switch s.currentContentNode.Kind() {
	case node.Map:
		s.nextByteExpecter = mapNodeByteExpecter{}
	case node.Array:
		s.nextByteExpecter = arrayNodeByteExpecter{}
	}
}

func (s *scanner) appendValueNode(nd node.Node) error {
	if s.node == nil {
		s.node = nd

		return nil
	}

	err := s.validateAppendValueNode(nd)
	if err != nil {
		return err
	}

	s.appendNode(nd)

	return nil
}

func (s *scanner) appendNode(nd node.Node) {
	node.UnsafeAppend(s.currentContentNode, nd)
}

func (s *scanner) validateAppendValueNode(nd node.Node) error {
	if s.currentContentNode == nil {
		return ErrContentableNodeIsNotInitialized
	}

	if s.currentContentNode.Kind() == node.Map {
		err := s.validateAppendInMapNode(nd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *scanner) validateAppendInMapNode(nd node.Node) error {
	if s.currentContentNode == nil {
		return ErrContentableNodeIsNotInitialized
	}

	if len(s.currentContentNode.Content())%2 == 0 {
		if nd.Kind() != node.String {
			return ErrMissingProperty
		}
	}

	return nil
}

func (s *scanner) closeArrayNode() error {
	if s.currentContentNode == nil {
		return ErrContentableNodeIsNotInitialized
	}

	if s.currentContentNode.Kind() != node.Array {
		return ErrWrongCloseArrayNode
	}

	return s.closeContentableNode()
}

func (s *scanner) closeMapNode() error {
	if s.currentContentNode.Kind() != node.Map {
		return ErrWrongCloseMapNode
	}

	if len(s.currentContentNode.Content())%2 == 1 {
		return ErrMissingPropertyValue
	}

	return s.closeContentableNode()
}

func (s *scanner) closeContentableNode() error {
	notClosedContentableNodesCount := len(s.contentableNodes)

	switch notClosedContentableNodesCount {
	case 0:
		return ErrNoneContentableNodeLeft
	case 1:
		s.currentContentNode = nil
		s.contentableNodes = nil
	default:
		prevContentNode := s.contentableNodes[notClosedContentableNodesCount-2]
		s.contentableNodes = s.contentableNodes[:notClosedContentableNodesCount-1]

		s.currentContentNode = prevContentNode

		s.setCloseContentableNodeNextByteExpecter()
	}

	return nil
}

func (s *scanner) setCloseContentableNodeNextByteExpecter() {
	switch s.currentContentNode.Kind() {
	case node.Map:
		s.nextByteExpecter = mapValueNodeExpecter{}
	case node.Array:
		s.nextByteExpecter = arrayValueNodeExpecter{}
	}
}

func (s *scanner) allContentableNodesClosed() bool {
	return len(s.contentableNodes) == 0
}
