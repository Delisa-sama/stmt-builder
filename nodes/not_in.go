package nodes

// NotInNode represents set operator
type NotInNode struct {
	name  NameNode
	right ArrayNode
}

// NewNotInNode returns NotInNode
func NewNotInNode(name NameNode, right ArrayNode) NotInNode {
	return NotInNode{
		name:  name,
		right: right,
	}
}

// Name returns name
func (n NotInNode) Name() NameNode {
	return n.name
}

// Childs returns collection of childs
func (n NotInNode) Childs() []Node {
	return []Node{n.right}
}

// Accept accepts translate visitor to invoke TranslateNotInNode method
func (n NotInNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateNotInNode(n)
}
