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
		return nodes.NewNullNode()
	}
	switch reflect.TypeOf(operand).Kind() {
	case reflect.Int:
		return nodes.NewIntNode(int64(operand.(int)))
	case reflect.Int8:
		return nodes.NewIntNode(int64(operand.(int8)))
	case reflect.Int16:
		return nodes.NewIntNode(int64(operand.(int16)))
	case reflect.Int32:
		return nodes.NewIntNode(int64(operand.(int32)))
	case reflect.Int64:
		return nodes.NewIntNode(operand.(int64))
	case reflect.Uint:
		return nodes.NewUintNode(uint64(operand.(uint)))
	case reflect.Uint8:
		return nodes.NewUintNode(uint64(operand.(uint8)))
	case reflect.Uint16:
		return nodes.NewUintNode(uint64(operand.(uint16)))
	case reflect.Uint32:
		return nodes.NewUintNode(uint64(operand.(uint32)))
	case reflect.Uint64:
		return nodes.NewUintNode(operand.(uint64))
	case reflect.Float64:
		return nodes.NewFloatNode(operand.(float64))
	case reflect.Float32:
		return nodes.NewFloatNode(float64(operand.(float32)))
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
			cast = append(cast, castToNode(s.Index(i).Interface()))
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
