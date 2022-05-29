package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// LT represents larger than operator
type LT struct{}

// Node returns LtNode
func (o LT) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewLtNode(nodes.NewNameNode(leftOp), rightOp)
}
