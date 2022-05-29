package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// GE represents greater or equals operator
type GE struct{}

// Node returns GENode
func (o GE) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewGENode(nodes.NewNameNode(leftOp), rightOp)
}
