package node

type IterationStep interface {
	KeyValue(key, value Node) (resultKey, resultValue Node, err error)
}

type MapSource interface {
	MapNode() (MapNode, error)
}

type iterationMapSource struct {
	source   MapSource
	iterStep IterationStep
}

func IterationMapSource(source MapSource, iterStep IterationStep) MapSource {
	return iterationMapSource{
		source:   source,
		iterStep: iterStep,
	}
}

func (m iterationMapSource) MapNode() (MapNode, error) {
	mapNode, err := m.source.MapNode()
	if err != nil {
		return mapNode, err
	}

	iter := MapNodeIterator(mapNode.Content())

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

type joinIterationStep struct {
	steps []IterationStep
}

func JoinIterationSteps(steps ...IterationStep) IterationStep {
	return joinIterationStep{
		steps: steps,
	}
}

func (i joinIterationStep) KeyValue(key, value Node) (resKey, resValue Node, err error) {
	resKey, resValue = key, value

	for _, step := range i.steps {
		resKey, resValue, err = step.KeyValue(resKey, resValue)
		if err != nil {
			return nil, nil, err
		}
	}

	return resKey, resValue, nil
}
