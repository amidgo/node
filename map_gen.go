package node

type IterationStep interface {
	KeyValue(key, value Node) (resultKey, resultValue Node, err error)
}

type MapSource interface {
	MapNode() (MapNode, error)
}

type IterationMapSource struct {
	source   MapSource
	iterStep IterationStep
}

func NewIterationMapSource(source MapSource, iterStep IterationStep) *IterationMapSource {
	return &IterationMapSource{
		source:   source,
		iterStep: iterStep,
	}
}

func (m *IterationMapSource) MapNode() (MapNode, error) {
	mapNode, err := m.source.MapNode()
	if err != nil {
		return mapNode, err
	}

	iter := MakeMapNodeIterator(mapNode.Content())

	content := make([]Node, 0, len(mapNode.Content()))

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
