package nodes

type NeNode struct {
	Left  Node
	Right Node
}

func (n NeNode) Childs() []Node {
	return []Node{n.Left, n.Right}
}
