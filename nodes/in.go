package nodes

type InNode struct {
	Left  Node
	Right Node
}

func (n InNode) Childs() []Node {
	return []Node{n.Left, n.Right}
}
