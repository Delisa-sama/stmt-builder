package statement

import (
	"github.com/Delisa-sama/stmt-builder/limit"
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/offset"
	"github.com/Delisa-sama/stmt-builder/sort"
)

// Statement represents tree of nodes that can be translated to query
type Statement struct {
	root   nodes.Node
	sort   sort.Sort
	limit  limit.Limit
	offset offset.Offset
}

// Operator represents operator for construct statement
type Operator interface {
	Node(leftOp string) nodes.Node
}

// New returns new statement
func New(leftOperand string, op Operator) Statement {
	return Statement{
		root: op.Node(leftOperand),
		sort: nil,
	}
}

// Empty returns new empty statement
func Empty() Statement {
	return Statement{
		root: nil,
	}
}

// GetRoot returns root node
func (s Statement) GetRoot() nodes.Node {
	return s.root
}

// GetSort returns sort
func (s Statement) GetSort() sort.Sort {
	return s.sort
}

// GetLimit returns the limit
func (s Statement) GetLimit() limit.Limit {
	return s.limit
}

// GetOffset returns the offset
func (s Statement) GetOffset() offset.Offset {
	return s.offset
}

// IsEmpty returns true if statement is empty
func (s Statement) IsEmpty() bool {
	return s.root == nil
}

// Not returns new statement with negotiate
func Not(statement Statement) Statement {
	return Statement{
		root: nodes.NewNotNode(statement.GetRoot()),
	}
}

// And returns a new statement concatenating the two with AND operator
func (s Statement) And(another Statement) Statement {
	if s.IsEmpty() {
		return Statement{root: another.GetRoot()}
	}
	return Statement{
		root: nodes.NewAndNode(s.GetRoot(), another.GetRoot()),
	}
}

// Or returns a new statement concatenating the two with OR operator
func (s Statement) Or(another Statement) Statement {
	if s.IsEmpty() {
		return Statement{root: another.GetRoot()}
	}
	return Statement{
		root: nodes.NewOrNode(s.GetRoot(), another.GetRoot()),
	}
}

// Sort returns statement with sort
func (s Statement) Sort(columnNames []string, direction sort.Direction) Statement {
	if len(columnNames) == 0 {
		return s
	}
	s.sort = sort.NewSort(columnNames, direction)
	return s
}

// Limit returns statement with the limit
func (s Statement) Limit(l uint) Statement {
	s.limit = limit.NewLimit(l)
	return s
}

// Offset returns statement with the offset
func (s Statement) Offset(l uint) Statement {
	s.offset = offset.NewOffset(l)
	return s
}
