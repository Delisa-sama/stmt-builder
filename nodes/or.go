package nodes

type OrNode struct {
	Left  Node
	Right Node
}

func (n OrNode) Childs() []Node {
	return []Node{n.Left, n.Right}
}
