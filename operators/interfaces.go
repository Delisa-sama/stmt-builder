package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// SingleValue represents abstract value that can returns equivalent Node
type SingleValue interface {
	Node() nodes.Node
}

// MultipleValues represents abstract values that can returns equivalent Node collection of Node
type MultipleValues interface {
	Node() []nodes.Node
}
