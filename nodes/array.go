package nodes

type ArrayNode struct {
	Value []Node
}

func (n ArrayNode) Childs() []Node {
	return n.Value
}
