package query

import (
	"reflect"
	"time"

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

func (s Statement) Root() nodes.Node {
	return s.root
}

func Not(statement Statement) Statement {
	return Statement{
		root: nodes.NewNotNode(statement.root),
	}
}

func castToNode(operand any) nodes.Node {
	if operand == nil {
		return nodes.NullNode{}
	}
	switch reflect.TypeOf(operand).Kind() {
	case reflect.String:
		return nodes.NewStringNode(operand.(string))
	case reflect.Struct:
		switch typedOperand := operand.(type) {
		case time.Time:
			return nodes.NewTimeNode(typedOperand)
		default:
			return nodes.NewValueNode(typedOperand)
		}
	case reflect.Slice:
		s := reflect.ValueOf(operand)
		cast := make([]nodes.Node, 0, s.Len())
		for i := 0; i < s.Len(); i++ {
			var castNode nodes.Node
			kind := s.Index(i).Kind()
			if kind == reflect.String {
				castNode = nodes.NewStringNode(s.Index(i).String())
			} else {
				castNode = nodes.NewValueNode(s.Index(i).Interface())
			}
			cast = append(cast, castNode)
		}
		return nodes.NewArrayNode(cast)
	default:
		return nodes.NewValueNode(operand)
	}
}

func (s Statement) And(another Statement) Statement {
	return Statement{
		root: nodes.NewAndNode(s.root, another.root),
	}
}

func (s Statement) Or(another Statement) Statement {
	return Statement{
		root: nodes.NewOrNode(s.root, another.root),
	}
}
