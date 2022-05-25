package nodes

type EqNode struct {
	Left  Node
	Right Node
}

func (n EqNode) Childs() []Node {
	return []Node{n.Left, n.Right}
}
