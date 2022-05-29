package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// NE represents not equals operator
type NE struct{}

// Node returns NeNode
func (o NE) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewNeNode(nodes.NewNameNode(leftOp), rightOp)
}
