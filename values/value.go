package values

import (
	"reflect"
	"time"

	"github.com/Delisa-sama/stmt-builder/nodes"
)

// Value represents value to construct statement
type Value struct {
	typeOf Type
	value  interface{}
}

// Type represents type of value
type Type uint8

// currently supported types
const (
	UnknownType Type = iota
	NullType
	BoolType
	IntType
	Int64Type
	Uint64Type
	Float64Type
	StringType
	TimeType
)

// Node returns node based on value
func (v Value) Node() nodes.Node {
	return castToNode(v.typeOf, v.value)
}

// Values represents multiple values to construct a statement
type Values struct {
	typeOf Type
	value  interface{}
}

// Node returns collection of nodes from values
func (v Values) Node() []nodes.Node {
	s := reflect.ValueOf(v.value)
	cast := make([]nodes.Node, 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		cast = append(cast, castToNode(v.typeOf, s.Index(i).Interface()))
	}
	return cast
}

// castToNode returns new node by value and type
func castToNode(typeOf Type, value interface{}) nodes.Node {
	switch typeOf {
	case NullType:
		return nodes.NewNullNode()
	case BoolType:
		return nodes.NewBoolNode(value.(bool))
	case IntType:
		return nodes.NewIntNode(int64(value.(int)))
	case Int64Type:
		return nodes.NewIntNode(value.(int64))
	case Float64Type:
		return nodes.NewFloatNode(value.(float64))
	case Uint64Type:
		return nodes.NewUintNode(value.(uint64))
	case StringType:
		return nodes.NewStringNode(value.(string))
	case TimeType:
		return nodes.NewTimeNode(value.(time.Time))
	}

	return nodes.NullNode{}
}

// Null returns new Null typed value
func Null() Value {
	return Value{
		typeOf: NullType,
		value:  nil,
	}
}

// Bool returns new Bool typed value.
func Bool(v bool) Value {
	return Value{
		typeOf: BoolType,
		value:  v,
	}
}

// Bools returns new Bool typed multiple value
func Bools(v ...bool) Values {
	return Values{
		typeOf: BoolType,
		value:  v,
	}
}

// Int returns new int typed value.
func Int(v int) Value {
	return Value{
		typeOf: IntType,
		value:  v,
	}
}

// Ints returns new int typed multiple value
func Ints(v ...int) Values {
	return Values{
		typeOf: IntType,
		value:  v,
	}
}

// Int64 returns new int64 typed value.
func Int64(v int64) Value {
	return Value{
		typeOf: Int64Type,
		value:  v,
	}
}

// Ints64 returns new int64 typed multiple value
func Ints64(v ...int64) Values {
	return Values{
		typeOf: Int64Type,
		value:  v,
	}
}

// Uint returns new uint typed value.
func Uint(v uint64) Value {
	return Value{
		typeOf: Uint64Type,
		value:  v,
	}
}

// Uints returns new uint typed multiple value
func Uints(v ...uint64) Values {
	return Values{
		typeOf: Uint64Type,
		value:  v,
	}
}

// Float returns new float typed value.
func Float(v float64) Value {
	return Value{
		typeOf: Float64Type,
		value:  v,
	}
}

// Floats returns new float typed multiple value
func Floats(v ...float64) Values {
	return Values{
		typeOf: Float64Type,
		value:  v,
	}
}

// String returns new string typed value.
func String(v string) Value {
	return Value{
		typeOf: StringType,
		value:  v,
	}
}

// Strings returns new string typed multiple value
func Strings(v ...string) Values {
	return Values{
		typeOf: StringType,
		value:  v,
	}
}

// Time returns new time.Time typed value.
func Time(v time.Time) Value {
	return Value{
		typeOf: TimeType,
		value:  v,
	}
}

// Times returns new time.Time typed multiple value
func Times(v ...time.Time) Values {
	return Values{
		typeOf: TimeType,
		value:  v,
	}
}
