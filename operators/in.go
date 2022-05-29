package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// In represents set operator
type In struct{}

// Node returns InNode
func (o In) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewInNode(nodes.NewNameNode(leftOp), rightOp)
}
