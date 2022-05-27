package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// LEOperator represents larger or equals operator
type LEOperator struct{}

// Node returns LeNode
func (o LEOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewLeNode(nodes.NewNameNode(leftOp), rightOp)
}
