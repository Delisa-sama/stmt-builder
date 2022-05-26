package nodes

import (
	"fmt"
)

type Node interface {
	Childs() []Node
}

func NodeToSQL(node Node) string {
	switch typedNode := node.(type) {
	case ValueNode:
		return fmt.Sprintf("%v", typedNode.Value)
	case StringNode:
		return fmt.Sprintf("'%s'", typedNode.Value)
	case NameNode:
		return typedNode.Name
	case NullNode:
		return "NULL"
	case ArrayNode:
		return ","
	case EqNode:
		_, isNull := typedNode.Right.(NullNode)
		if isNull {
			return " IS "
		}
		return " = "
	case InNode:
		return " IN "
	case NeNode:
		_, isNull := typedNode.Right.(NullNode)
		if isNull {
			return " IS NOT "
		}
		return " <> "
	case AndNode:
		return " AND "
	case OrNode:
		return " OR "
	}

	return ""
}
