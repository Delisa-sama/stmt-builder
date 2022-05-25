package query

import (
	"strings"

	"github.com/Delisa-sama/stmt-builder/nodes"
)

type Statement struct {
	root nodes.Node
	args []nodes.Node
}

type Operator interface {
	Node(leftOp string, rightOp nodes.Node) nodes.Node
}

func NewStatement(leftOperand string, op Operator, rightOperand any) Statement {
	s := Statement{}
	var rightNode nodes.Node
	switch typedOperand := rightOperand.(type) {
	case nil:
		rightNode = nodes.NullNode{}
	case string:
		rightNode = nodes.StringNode{Value: typedOperand}
	default:
		rightNode = nodes.ValueNode{Value: rightOperand}
	}
	s.root = op.Node(leftOperand, rightNode)
	return s
}

func (s Statement) And(another Statement) Statement {
	return Statement{
		root: nodes.AndNode{
			Left:  s.root,
			Right: another.root,
		},
	}
}

func (s Statement) Or(another Statement) Statement {
	return Statement{
		root: nodes.OrNode{
			Left:  s.root,
			Right: another.root,
		},
	}
}

type Placeholder interface {
	Next() string
}

func (s Statement) ToSQL(placeholder Placeholder) (string, []interface{}) {
	query := s.toSQL(s.root, placeholder)
	var args []interface{}
	if placeholder != nil {
		nodeArgs := s.getArgs(s.root)
		for _, arg := range nodeArgs {
			args = append(args, arg)
		}
	}
	return query, args
}

const (
	openParentheses  = '('
	closeParentheses = ')'
	space            = ' '
)

func (s Statement) toSQL(node nodes.Node, placeholder Placeholder) string {
	if node == nil {
		return ""
	}

	parentheses := false
	switch node.(type) {
	case nodes.NameNode, nodes.NullNode:
		return nodes.NodeToSQL(node)
	case nodes.ValueNode, nodes.StringNode:
		if placeholder != nil {
			return placeholder.Next()
		}
		return nodes.NodeToSQL(node)
	case nodes.AndNode, nodes.OrNode:
		parentheses = true
	}

	childs := node.Childs()
	// non leaf nodes without childs are useless, ignore it
	if len(childs) == 0 {
		return ""
	}

	queryBuilder := strings.Builder{}
	if parentheses {
		queryBuilder.WriteRune(openParentheses)
	}
	queryBuilder.WriteString(s.toSQL(childs[0], placeholder))
	for _, child := range childs[1:] {
		queryBuilder.WriteRune(space)
		queryBuilder.WriteString(nodes.NodeToSQL(node))
		queryBuilder.WriteRune(space)
		queryBuilder.WriteString(s.toSQL(child, placeholder))
	}
	if parentheses {
		queryBuilder.WriteRune(closeParentheses)
	}

	return queryBuilder.String()
}

func (s Statement) getArgs(node nodes.Node) []nodes.Node {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case nodes.ValueNode, nodes.StringNode:
		return []nodes.Node{node}
	}
	childs := node.Childs()
	// non leaf nodes without childs are useless, ignore it
	if len(childs) == 0 {
		return nil
	}
	args := make([]nodes.Node, 0)
	for _, child := range childs {
		args = append(args, s.getArgs(child)...)
	}

	return args
}
