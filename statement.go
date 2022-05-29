package query

import (
	"reflect"
	"time"

	"github.com/Delisa-sama/stmt-builder/nodes"
)

// Statement represents tree of nodes that can be translated to query
type Statement struct {
	root nodes.Node
}

// Operator represents operator for construct statement
type Operator interface {
	Node(leftOp string, rightOp nodes.Node) nodes.Node
}

// NewStatement returns new statement
func NewStatement(leftOperand string, op Operator, rightOperand interface{}) Statement {
	return Statement{
		root: op.Node(leftOperand, castToNode(rightOperand)),
	}
}

// NewEmptyStatement returns new empty statement
func NewEmptyStatement() Statement {
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

func castToNode(operand interface{}) nodes.Node {
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
