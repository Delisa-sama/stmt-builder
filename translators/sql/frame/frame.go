package frame

import (
	"strings"

	"github.com/Delisa-sama/stmt-builder/nodes"
)

// Frame represents recursion iteration for depth-first traversal algorithm.
type Frame struct {
	node             nodes.Node
	parent           *Frame
	returnValues     map[int]string
	returnValueIndex int
	localState       local
	execCount        int
}

type local struct {
	statementParentheses bool
	childsParentheses    bool
	queryBuilder         strings.Builder
}

// NewFrame return new instance of Frame
func NewFrame(node nodes.Node, parent *Frame, index int) *Frame {
	return &Frame{
		node:             node,
		parent:           parent,
		returnValueIndex: index,
		returnValues:     make(map[int]string),
	}
}

// Result returns result of iteration
func (f *Frame) Result() string {
	return f.localState.queryBuilder.String()
}

// PassToParent passes return value to parent frame
func (f *Frame) PassToParent(returnValue string) {
	if f.parent == nil {
		return
	}

	f.parent.returnValues[f.returnValueIndex] = returnValue
}

const (
	openParentheses  = '('
	closeParentheses = ')'
)

// TranslateFrame translates frame to SQL.
func (f *Frame) TranslateFrame(t nodes.TranslateVisitor) string {
	if f.node == nil {
		return ""
	}

	switch f.node.(type) {
	case nodes.AndNode, nodes.OrNode:
		f.localState.statementParentheses = true
	case nodes.InNode, nodes.NotNode, nodes.NotInNode:
		f.localState.childsParentheses = true
	}

	if nodeWithValue, ok := f.node.(nodes.NodeWithValue); ok {
		return nodeWithValue.Accept(t)
	}

	nodeWithChilds, ok := f.node.(nodes.NodeWithChilds)
	if !ok {
		return f.node.Accept(t)
	}

	childs := nodeWithChilds.Childs()
	// non leaf nodes without childs are useless, ignore it
	if len(childs) == 0 {
		return ""
	}

	if f.localState.statementParentheses {
		f.localState.queryBuilder.WriteRune(openParentheses)
	}
	if nodeWithName, ok := f.node.(nodes.NodeWithName); ok {
		name := nodeWithName.Name()
		f.localState.queryBuilder.WriteString(name.Accept(t))
	}
	// unary op
	if len(childs) == 1 {
		// TODO: think about exclusive array node translation
		_, isArrayNode := f.node.(nodes.ArrayNode)
		if !isArrayNode {
			f.localState.queryBuilder.WriteString(f.node.Accept(t))
		}
		if f.localState.childsParentheses {
			f.localState.queryBuilder.WriteRune(openParentheses)
		}
		f.localState.queryBuilder.WriteString(f.returnValues[0])
		if f.localState.childsParentheses {
			f.localState.queryBuilder.WriteRune(closeParentheses)
		}
	}
	// binary op
	if len(childs) == 2 {
		f.localState.queryBuilder.WriteString(f.returnValues[0])
		f.localState.queryBuilder.WriteString(f.node.Accept(t))
		f.localState.queryBuilder.WriteString(f.returnValues[1])
	}
	// variadic op
	if len(childs) > 2 {
		f.localState.queryBuilder.WriteString(f.returnValues[0])
		for i := range childs[1:] {
			f.localState.queryBuilder.WriteString(f.node.Accept(t))
			f.localState.queryBuilder.WriteString(f.returnValues[i])
		}
	}
	if f.localState.statementParentheses {
		f.localState.queryBuilder.WriteRune(closeParentheses)
	}

	return f.Result()

}

// GetNextChild returns next child for traversing.
// Returns nil if frame has no child.
func (f *Frame) GetNextChild() *Frame {
	f.execCount++
	if f.node == nil {
		return nil
	}

	nodeWithChilds, ok := f.node.(nodes.NodeWithChilds)
	if !ok {
		return nil
	}
	childs := nodeWithChilds.Childs()
	// non leaf nodes without childs are useless, ignore it
	if len(childs) == 0 {
		return nil
	}

	// unary op
	if len(childs) == 1 && f.execCount < 2 {
		return NewFrame(childs[0], f, 0)
	}
	// binary op
	if len(childs) == 2 && f.execCount < 3 {
		if f.execCount == 1 {
			return NewFrame(childs[0], f, 0)
		}
		return NewFrame(childs[1], f, 1)
	}
	// variadic op
	if len(childs) > 2 && f.execCount < len(childs)+1 {
		if f.execCount == 1 {
			return NewFrame(childs[0], f, 0)
		}
		return NewFrame(childs[f.execCount-1], f, 1)
	}

	return nil
}
