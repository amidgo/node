package json

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/amidgo/node"
)

var (
	ErrInvalidMapKeyKind     = errors.New("invalid map key kind")
	ErrInvalidBoolNode       = errors.New("invalid bool node")
	ErrInvalidIntegerNode    = errors.New("invalid integer node")
	ErrInvalidFloatNode      = errors.New("invalid float node")
	ErrInvalidValueNodeKind  = errors.New("invalid value node kind")
	ErrInvalidMapContentSize = errors.New("invalid map content size")
)

type nodeWriter struct {
	output io.Writer
}

func makeNodeWriterWithOutput(output io.Writer) *nodeWriter {
	return &nodeWriter{
		output: output,
	}
}

func (w *nodeWriter) writeNode(nd node.Node) error {
	switch nd.Type() {
	case node.Content:
		return w.writeContentNode(nd)
	case node.Value:
		return w.writeValue(nd)
	}

	return nil
}

func (w *nodeWriter) writeContentNode(nd node.Node) error {
	switch nd.Kind() {
	case node.Array:
		return w.writeArray(nd)
	case node.Map:
		return w.writeMap(nd)
	}

	return nil
}

func (w *nodeWriter) writeArray(nd node.Node) error {
	err := w.beginArray()
	if err != nil {
		return err
	}

	err = w.writeArrayContent(nd.Content())
	if err != nil {
		return err
	}

	return w.endArray()
}

func (w *nodeWriter) beginArray() error {
	err := w.writeByte('[')
	if err != nil {
		return err
	}

	return nil
}

func (w *nodeWriter) endArray() error {
	err := w.writeByte(']')
	if err != nil {
		return err
	}

	return nil
}

func (w *nodeWriter) writeArrayContent(content []node.Node) error {
	if len(content) == 0 {
		return nil
	}

	err := w.writeNode(content[0])
	if err != nil {
		return err
	}

	for _, nd := range content[1:] {
		err := w.writeByte(',')
		if err != nil {
			return err
		}

		err = w.writeNode(nd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *nodeWriter) writeMap(nd node.Node) error {
	err := w.beginMap()
	if err != nil {
		return err
	}

	err = w.writeMapContent(nd.Content())
	if err != nil {
		return err
	}

	return w.endMap()
}

func (w *nodeWriter) beginMap() error {
	err := w.writeByte('{')
	if err != nil {
		return err
	}

	return nil
}

func (w *nodeWriter) endMap() error {
	err := w.writeByte('}')
	if err != nil {
		return err
	}

	return nil
}

func (w *nodeWriter) writeMapContent(content []node.Node) error {
	if len(content) == 0 {
		return nil
	}

	if len(content)%2 == 1 {
		return ErrInvalidMapContentSize
	}

	iterator := node.MapNodeIterator(content)

	if iterator.HasNext() {
		err := w.writeNextMapIteratorNodePair(iterator)
		if err != nil {
			return err
		}
	}

	for iterator.HasNext() {
		err := w.writeByte(',')
		if err != nil {
			return err
		}

		err = w.writeNextMapIteratorNodePair(iterator)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *nodeWriter) writeNextMapIteratorNodePair(iterator node.Iterator) error {
	key, value := iterator.Next()
	if key.Kind() != node.String {
		return ErrInvalidMapKeyKind
	}

	err := w.writeString(key)
	if err != nil {
		return err
	}

	err = w.writeByte(':')
	if err != nil {
		return err
	}

	return w.writeNode(value)
}

func (w *nodeWriter) writeValue(nd node.Node) error {
	switch nd.Kind() {
	case node.Bool:
		return w.writeBool(nd)
	case node.Float:
		return w.writeFloat(nd)
	case node.Integer:
		return w.writeInteger(nd)
	case node.Empty:
		return w.writeEmpty()
	case node.String:
		return w.writeString(nd)
	default:
		return ErrInvalidValueNodeKind
	}
}

func (w *nodeWriter) writeBool(nd node.Node) error {
	boolNode, valid := nd.(interface{ Bool() bool })
	if !valid {
		return ErrInvalidBoolNode
	}

	if boolNode.Bool() {
		err := w.writeData([]byte{'t', 'r', 'u', 'e'})
		if err != nil {
			return err
		}
	} else {
		err := w.writeData([]byte{'f', 'a', 'l', 's', 'e'})
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *nodeWriter) writeInteger(nd node.Node) error {
	integerNode, valid := nd.(interface{ Int() int64 })
	if !valid {
		return ErrInvalidIntegerNode
	}

	s := strconv.FormatInt(integerNode.Int(), 10)

	err := w.writeStringData(s)
	if err != nil {
		return err
	}

	return nil
}

func (w *nodeWriter) writeFloat(nd node.Node) error {
	floatNode, valid := nd.(interface{ Float() float64 })
	if !valid {
		return ErrInvalidFloatNode
	}

	s := strconv.FormatFloat(floatNode.Float(), 'g', -1, 64)

	err := w.writeStringData(s)
	if err != nil {
		return err
	}

	return nil
}

func (w *nodeWriter) writeEmpty() error {
	err := w.writeData([]byte{'n', 'u', 'l', 'l'})
	if err != nil {
		return err
	}

	return nil
}

func (w *nodeWriter) writeString(nd node.Node) error {
	err := w.writeByte('"')
	if err != nil {
		return err
	}

	err = w.writeStringData(nd.Value())
	if err != nil {
		return fmt.Errorf("failed write value to output, %w", err)
	}

	err = w.writeByte('"')
	if err != nil {
		return err
	}

	return nil
}

func (w *nodeWriter) writeByte(b byte) error {
	_, err := w.output.Write([]byte{b})
	if err != nil {
		return fmt.Errorf("failed write '%s' to output, %w", string(b), err)
	}

	return nil
}

func (w *nodeWriter) writeData(data []byte) error {
	_, err := w.output.Write(data)
	if err != nil {
		return fmt.Errorf("failed write '%s' to output, %w", string(data), err)
	}

	return nil
}

func (w *nodeWriter) writeStringData(s string) error {
	_, err := io.WriteString(w.output, s)
	if err != nil {
		return fmt.Errorf("failed write '%s' to output, %w", s, err)
	}

	return nil
}
