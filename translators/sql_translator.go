package translators

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	query "github.com/Delisa-sama/stmt-builder"
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// Placeholder represents abstract placeholder for SQL query
type Placeholder interface {
	Next() string
}

// SQLTranslator represents translator from statement to SQL
type SQLTranslator struct {
	placeholder Placeholder
}

// SQLTranslatorOption represents option for SQLTranslator
type SQLTranslatorOption func(translator *SQLTranslator)

// WithPlaceholder option to define placeholder for SQLTranslator
func WithPlaceholder(placeholder Placeholder) SQLTranslatorOption {
	return func(translator *SQLTranslator) {
		translator.placeholder = placeholder
	}
}

// NewSQLTranslator returns new SQLTranslator
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

// GetArgs returns args for SQL query from statement
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
	nodeWithChilds, ok := node.(nodes.NodeWithChilds)
	if !ok {
		return []nodes.Node{node}
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

// Translate translates statement to SQL
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

// TranslateAndNode translates and node to sql
func (t *SQLTranslator) TranslateAndNode(node nodes.AndNode) string {
	return " AND "
}

// TranslateArrayNode translates array node to sql
func (t *SQLTranslator) TranslateArrayNode(node nodes.ArrayNode) string {
	return ","
}

// TranslateEqNode translates EQ node to sql
func (t *SQLTranslator) TranslateEqNode(node nodes.EqNode) string {
	_, isNull := node.Right().(nodes.NullNode)
	if isNull {
		return " IS "
	}
	return " = "
}

// TranslateGeNode translates GE node to sql
func (t *SQLTranslator) TranslateGeNode(node nodes.GeNode) string {
	return " >= "
}

// TranslateGtNode translates GT node to sql
func (t *SQLTranslator) TranslateGtNode(node nodes.GtNode) string {
	return " > "
}

// TranslateInNode translates In node to sql
func (t *SQLTranslator) TranslateInNode(node nodes.InNode) string {
	return " IN "
}

// TranslateLeNode translates LE node to sql
func (t *SQLTranslator) TranslateLeNode(node nodes.LeNode) string {
	return " <= "
}

// TranslateLtNode translates LT node to sql
func (t *SQLTranslator) TranslateLtNode(node nodes.LtNode) string {
	return " < "
}

// TranslateNameNode translates name node to sql
func (t *SQLTranslator) TranslateNameNode(node nodes.NameNode) string {
	return node.Name()
}

// TranslateNeNode translates ne node to sql
func (t *SQLTranslator) TranslateNeNode(node nodes.NeNode) string {
	_, isNull := node.Right().(nodes.NullNode)
	if isNull {
		return " IS NOT "
	}
	return " <> "
}

// TranslateNotNode translates not node to sql
func (t *SQLTranslator) TranslateNotNode(node nodes.NotNode) string {
	return "!"
}

// TranslateNullNode translates null node to sql
func (t *SQLTranslator) TranslateNullNode(node nodes.NullNode) string {
	return "NULL"
}

// TranslateOrNode translates or node to sql
func (t *SQLTranslator) TranslateOrNode(node nodes.OrNode) string {
	return " OR "
}

// TranslateStringNode translates string node to sql
func (t *SQLTranslator) TranslateStringNode(node nodes.StringNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return fmt.Sprintf("'%s'", node)
}

// TranslateValueNode translates value node to sql
func (t *SQLTranslator) TranslateValueNode(node nodes.ValueNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return fmt.Sprintf("%v", node.Value())
}

// TranslateTimeNode translates time node to sql
func (t *SQLTranslator) TranslateTimeNode(node nodes.TimeNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return fmt.Sprintf("'%s'", node.Format(time.RFC3339))
}

// TranslateIntNode translates int node to sql
func (t *SQLTranslator) TranslateIntNode(n nodes.IntNode) string {
	return strconv.Itoa(int(n))
}

// TranslateFloatNode translates float node to sql
func (t *SQLTranslator) TranslateFloatNode(n nodes.FloatNode) string {
	return strconv.FormatFloat(float64(n), 'f', -1, 64)
}

// TranslateUintNode translates uint node to SQL
func (t *SQLTranslator) TranslateUintNode(n nodes.UintNode) string {
	return strconv.FormatUint(uint64(n), 10)
}
