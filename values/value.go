package values

import (
	"reflect"
	"time"

	"github.com/Delisa-sama/stmt-builder/nodes"
)

// Value represents value to construct statement
type Value struct {
	typeOf  Type
	isSlice bool
	value   interface{}
}

// Type represents type of value
type Type uint8

// currently supported types
const (
	UnknownType Type = iota
	NullType
	BoolType
	Int64Type
	Uint64Type
	Float64Type
	StringType
	TimeType
)

// Node returns node based on value
func (v Value) Node() nodes.Node {
	if v.isSlice {
		s := reflect.ValueOf(v.value)
		cast := make([]nodes.Node, 0, s.Len())
		for i := 0; i < s.Len(); i++ {
			cast = append(cast, castToNode(v.typeOf, s.Index(i).Interface()))
		}
		return nodes.NewArrayNode(cast)
	}
	return castToNode(v.typeOf, v.value)
}

// castToNode returns new node by value and type
func castToNode(typeOf Type, value interface{}) nodes.Node {
	switch typeOf {
	case NullType:
		return nodes.NewNullNode()
	case BoolType:
		return nodes.NewValueNode(value)
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
		typeOf:  NullType,
		isSlice: false,
		value:   nil,
	}
}

// Bool returns new Bool typed value.
// Can pass multiple values.
func Bool(v ...bool) Value {
	vc := len(v)
	if vc == 0 {
		return Null()
	}
	if vc == 1 {
		return Value{
			typeOf:  BoolType,
			isSlice: false,
			value:   v[0],
		}
	}
	return Value{
		typeOf:  BoolType,
		isSlice: true,
		value:   v,
	}
}

// Int returns new int typed value.
// Can pass multiple values.
func Int(v ...int64) Value {
	vc := len(v)
	if vc == 0 {
		return Null()
	}
	if vc == 1 {
		return Value{
			typeOf:  Int64Type,
			isSlice: false,
			value:   v[0],
		}
	}
	return Value{
		typeOf:  Int64Type,
		isSlice: true,
		value:   v,
	}
}

// Uint returns new uint typed value.
// Can pass multiple values.
func Uint(v ...uint64) Value {
	vc := len(v)
	if vc == 0 {
		return Null()
	}
	if vc == 1 {
		return Value{
			typeOf:  Uint64Type,
			isSlice: false,
			value:   v[0],
		}
	}
	return Value{
		typeOf:  Uint64Type,
		isSlice: true,
		value:   v,
	}
}

// Float returns new float typed value.
// Can pass multiple values.
func Float(v ...float64) Value {
	vc := len(v)
	if vc == 0 {
		return Null()
	}
	if vc == 1 {
		return Value{
			typeOf:  Float64Type,
			isSlice: false,
			value:   v[0],
		}
	}
	return Value{
		typeOf:  Float64Type,
		isSlice: true,
		value:   v,
	}
}

// String returns new string typed value.
// Can pass multiple values.
func String(v ...string) Value {
	vc := len(v)
	if vc == 0 {
		return Null()
	}
	if vc == 1 {
		return Value{
			typeOf:  StringType,
			isSlice: false,
			value:   v[0],
		}
	}
	return Value{
		typeOf:  StringType,
		isSlice: true,
		value:   v,
	}
}

// Time returns new time.Time typed value.
// Can pass multiple values.
func Time(v ...time.Time) Value {
	vc := len(v)
	if vc == 0 {
		return Null()
	}
	if vc == 1 {
		return Value{
			typeOf:  TimeType,
			isSlice: false,
			value:   v[0],
		}
	}
	return Value{
		typeOf:  TimeType,
		isSlice: true,
		value:   v,
	}
}
