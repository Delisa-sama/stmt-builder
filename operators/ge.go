package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// GEOperator represents greater or equals operator
type GEOperator struct{}

// Node returns GENode
func (o GEOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewGENode(nodes.NewNameNode(leftOp), rightOp)
}
