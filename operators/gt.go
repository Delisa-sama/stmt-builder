package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// GT represents greater than operator
type GT struct{}

// Node returns GtNode
func (o GT) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewGtNode(nodes.NewNameNode(leftOp), rightOp)
}
