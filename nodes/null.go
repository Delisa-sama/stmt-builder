package nodes

type NullNode struct{}

func (n NullNode) Childs() []Node {
	return nil
}
