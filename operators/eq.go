package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// EQ represents equals operator
type EQ struct{}

// Node returns EqNode
func (o EQ) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewEqNode(nodes.NewNameNode(leftOp), rightOp)
}
