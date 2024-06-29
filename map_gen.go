package node

type IterationStep interface {
	KeyValue(key, value Node) (resultKey, resultValue Node, err error)
}

type MapIterableGenerate struct {
	source   MapNode
	iterStep IterationStep
}

func NewMapIterableGenerate(sourceNode MapNode, iterStep IterationStep) *MapIterableGenerate {
	return &MapIterableGenerate{
		source:   sourceNode,
		iterStep: iterStep,
	}
}

func (m *MapIterableGenerate) MapNode() (MapNode, error) {
	iter := MakeMapNodeIterator(m.source.Content())

	content := make([]Node, 0, len(m.source.Content()))

	for iter.HasNext() {
		key, value := iter.Next()

		resKey, resValue, err := m.iterStep.KeyValue(key, value)
		if err != nil {
			return MapNode{}, err
		}

		content = append(content, resKey, resValue)
	}

	return MakeMapNode(content...), nil
}

type JoinIterationStep struct {
	steps []IterationStep
}

func NewJoinIterationStep(steps ...IterationStep) *JoinIterationStep {
	return &JoinIterationStep{
		steps: steps,
	}
}

func (i *JoinIterationStep) KeyValue(key, value Node) (resKey, resValue Node, err error) {
	resKey, resValue = key, value

	for _, step := range i.steps {
		resKey, resValue, err = step.KeyValue(resKey, resValue)
		if err != nil {
			return nil, nil, err
		}
	}

	return resKey, resValue, nil
}
