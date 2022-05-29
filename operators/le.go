package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// LE represents larger or equals operator
type LE struct{}

// Node returns LeNode
func (o LE) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewLeNode(nodes.NewNameNode(leftOp), rightOp)
}
