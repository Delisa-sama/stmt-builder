package translators

import (
	"fmt"
	"strings"
	"time"

	query "github.com/Delisa-sama/stmt-builder"
	"github.com/Delisa-sama/stmt-builder/nodes"
)

type Placeholder interface {
	Next() string
}

type SQLTranslator struct {
	placeholder Placeholder
}

type SQLTranslatorOption func(translator *SQLTranslator)

func WithPlaceholder(placeholder Placeholder) SQLTranslatorOption {
	return func(translator *SQLTranslator) {
		translator.placeholder = placeholder
	}
}

func NewSQLTranslator(opts ...SQLTranslatorOption) *SQLTranslator {
	t := &SQLTranslator{}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

const (
	openParentheses  = '('
	closeParentheses = ')'
)

func (t *SQLTranslator) GetArgs(s query.Statement) []interface{} {
	statementArgs := t.getArgs(s.Root())
	if len(statementArgs) == 0 {
		return nil
	}
	args := make([]interface{}, 0, len(statementArgs))
	for _, arg := range statementArgs {
		args = append(args, arg)
	}
	return args
}

func (t *SQLTranslator) getArgs(node nodes.Node) []nodes.Node {
	if node == nil {
		return nil
	}
	switch node.(type) {
	case nodes.ValueNode, nodes.StringNode:
		return []nodes.Node{node}
	}
	nodeWithChilds, ok := node.(nodes.NodeWithChilds)
	if !ok {
		return nil
	}
	childs := nodeWithChilds.Childs()
	// non leaf nodes without childs are useless, ignore it
	if len(childs) == 0 {
		return nil
	}
	args := make([]nodes.Node, 0)
	for _, child := range childs {
		args = append(args, t.getArgs(child)...)
	}

	return args
}

func (t *SQLTranslator) Translate(s query.Statement) string {
	return t.translate(s.Root())
}

func (t *SQLTranslator) translate(node nodes.Node) string {
	if node == nil {
		return ""
	}

	statementParentheses := false
	childsParentheses := false
	switch node.(type) {
	case nodes.AndNode, nodes.OrNode:
		statementParentheses = true
	case nodes.InNode, nodes.NotNode:
		childsParentheses = true
	}

	nodeWithChilds, ok := node.(nodes.NodeWithChilds)
	if !ok {
		return node.Accept(t)
	}
	childs := nodeWithChilds.Childs()
	// non leaf nodes without childs are useless, ignore it
	if len(childs) == 0 {
		return ""
	}

	queryBuilder := strings.Builder{}
	if statementParentheses {
		queryBuilder.WriteRune(openParentheses)
	}
	if nodeWithName, ok := node.(nodes.NodeWithName); ok {
		name := nodeWithName.Name()
		queryBuilder.WriteString(name.Accept(t))
	}
	// unary op
	if len(childs) == 1 {
		queryBuilder.WriteString(node.Accept(t))
		if childsParentheses {
			queryBuilder.WriteRune(openParentheses)
		}
		queryBuilder.WriteString(t.translate(childs[0]))
		if childsParentheses {
			queryBuilder.WriteRune(closeParentheses)
		}
	}
	// binary op
	if len(childs) == 2 {
		queryBuilder.WriteString(t.translate(childs[0]))
		queryBuilder.WriteString(node.Accept(t))
		queryBuilder.WriteString(t.translate(childs[1]))
	}
	// variadic op
	if len(childs) > 2 {
		queryBuilder.WriteString(t.translate(childs[0]))
		for _, child := range childs[1:] {
			queryBuilder.WriteString(node.Accept(t))
			queryBuilder.WriteString(t.translate(child))
		}
	}
	if statementParentheses {
		queryBuilder.WriteRune(closeParentheses)
	}

	return queryBuilder.String()
}

func (t *SQLTranslator) TranslateAndNode(node nodes.AndNode) string {
	return " AND "
}

func (t *SQLTranslator) TranslateArrayNode(node nodes.ArrayNode) string {
	return ","
}

func (t *SQLTranslator) TranslateEqNode(node nodes.EqNode) string {
	_, isNull := node.Right().(nodes.NullNode)
	if isNull {
		return " IS "
	}
	return " = "
}

func (t *SQLTranslator) TranslateGeNode(node nodes.GeNode) string {
	return " >= "
}

func (t *SQLTranslator) TranslateGtNode(node nodes.GtNode) string {
	return " > "
}

func (t *SQLTranslator) TranslateInNode(node nodes.InNode) string {
	return " IN "
}

func (t *SQLTranslator) TranslateLeNode(node nodes.LeNode) string {
	return " <= "
}

func (t *SQLTranslator) TranslateLtNode(node nodes.LtNode) string {
	return " < "
}

func (t *SQLTranslator) TranslateNameNode(node nodes.NameNode) string {
	return node.Name()
}

func (t *SQLTranslator) TranslateNeNode(node nodes.NeNode) string {
	_, isNull := node.Right().(nodes.NullNode)
	if isNull {
		return " IS NOT "
	}
	return " <> "
}

func (t *SQLTranslator) TranslateNotNode(node nodes.NotNode) string {
	return "!"
}

func (t *SQLTranslator) TranslateNullNode(node nodes.NullNode) string {
	return "NULL"
}

func (t *SQLTranslator) TranslateOrNode(node nodes.OrNode) string {
	return " OR "
}

func (t *SQLTranslator) TranslateStringNode(node nodes.StringNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return fmt.Sprintf("'%s'", node.String())
}

func (t *SQLTranslator) TranslateValueNode(node nodes.ValueNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return fmt.Sprintf("%v", node.Value())
}

func (t *SQLTranslator) TranslateTimeNode(node nodes.TimeNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return fmt.Sprintf("'%s'", node.Time().Format(time.RFC3339))
}
