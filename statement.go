package query

import (
	"reflect"
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
	return Statement{
		root: op.Node(leftOperand, castToNode(rightOperand)),
	}
}

func castToNode(operand any) nodes.Node {
	if operand == nil {
		return nodes.NullNode{}
	}
	switch reflect.TypeOf(operand).Kind() {
	case reflect.String:
		return nodes.StringNode{Value: operand.(string)}
	case reflect.Slice:
		s := reflect.ValueOf(operand)
		cast := make([]nodes.Node, 0, s.Len())
		for i := 0; i < s.Len(); i++ {
			var castNode nodes.Node
			kind := s.Index(i).Kind()
			if kind == reflect.String {
				castNode = nodes.StringNode{Value: s.Index(i).String()}
			} else {
				castNode = nodes.ValueNode{Value: s.Index(i).Interface()}
			}
			cast = append(cast, castNode)
		}
		return nodes.ArrayNode{Value: cast}
	default:
		return nodes.ValueNode{Value: operand}
	}
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
)

func (s Statement) toSQL(node nodes.Node, placeholder Placeholder) string {
	if node == nil {
		return ""
	}

	statementParentheses := false
	rightChildsParentheses := false
	switch node.(type) {
	case nodes.NameNode, nodes.NullNode:
		return nodes.NodeToSQL(node)
	case nodes.ValueNode, nodes.StringNode:
		if placeholder != nil {
			return placeholder.Next()
		}
		return nodes.NodeToSQL(node)
	case nodes.AndNode, nodes.OrNode:
		statementParentheses = true
	case nodes.InNode:
		rightChildsParentheses = true
	case nodes.ArrayNode:
	}

	childs := node.Childs()
	// non leaf nodes without childs are useless, ignore it
	if len(childs) == 0 {
		return ""
	}

	queryBuilder := strings.Builder{}
	if statementParentheses {
		queryBuilder.WriteRune(openParentheses)
	}
	queryBuilder.WriteString(s.toSQL(childs[0], placeholder))
	for _, child := range childs[1:] {
		queryBuilder.WriteString(nodes.NodeToSQL(node))
		if rightChildsParentheses {
			queryBuilder.WriteRune(openParentheses)
		}
		queryBuilder.WriteString(s.toSQL(child, placeholder))
		if rightChildsParentheses {
			queryBuilder.WriteRune(closeParentheses)
		}
	}
	if statementParentheses {
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
