package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// GTOperator represents greater than operator
type GTOperator struct{}

// Node returns GtNode
func (o GTOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewGtNode(nodes.NewNameNode(leftOp), rightOp)
}
