package nodes

type NameNode struct {
	Name string
}

func (n NameNode) Childs() []Node {
	return nil
}
