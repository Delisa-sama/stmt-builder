package statement

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// Statement represents tree of nodes that can be translated to query
type Statement struct {
	root nodes.Node
}

// Operator represents operator for construct statement
type Operator interface {
	Node(leftOp string) nodes.Node
}

// New returns new statement
func New(leftOperand string, op Operator) Statement {
	return Statement{
		root: op.Node(leftOperand),
	}
}

// Empty returns new empty statement
func Empty() Statement {
	return Statement{
		root: nil,
	}
}

// Root returns root node
func (s Statement) Root() nodes.Node {
	return s.root
}

// IsEmpty returns true if statement is empty
func (s Statement) IsEmpty() bool {
	return s.root == nil
}

// Not returns new statement with negotiate
func Not(statement Statement) Statement {
	return Statement{
		root: nodes.NewNotNode(statement.Root()),
	}
}

// And returns a new statement concatenating the two with AND operator
func (s Statement) And(another Statement) Statement {
	if s.IsEmpty() {
		return Statement{root: another.Root()}
	}
	return Statement{
		root: nodes.NewAndNode(s.Root(), another.Root()),
	}
}

// Or returns a new statement concatenating the two with OR operator
func (s Statement) Or(another Statement) Statement {
	if s.IsEmpty() {
		return Statement{root: another.Root()}
	}
	return Statement{
		root: nodes.NewOrNode(s.Root(), another.Root()),
	}
}
