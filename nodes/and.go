package nodes

type AndNode struct {
	Left  Node
	Right Node
}

func (n AndNode) Childs() []Node {
	return []Node{n.Left, n.Right}
}
