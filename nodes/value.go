package nodes

type ValueNode struct {
	Value any
}

func (n ValueNode) Childs() []Node {
	return nil
}
