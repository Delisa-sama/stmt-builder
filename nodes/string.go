package nodes

type StringNode struct {
	Value string
}

func (n StringNode) Childs() []Node {
	return nil
}
