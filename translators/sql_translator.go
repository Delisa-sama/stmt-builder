package translators

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/sort"
	"github.com/Delisa-sama/stmt-builder/statement"
)

// Placeholder represents abstract placeholder for SQL query
type Placeholder interface {
	Next() string
}

// SQLTranslator represents translator from statement to SQL
type SQLTranslator struct {
	placeholder Placeholder
}

// Statement represents abstract statement
type Statement interface {
	GetRoot() nodes.Node
	GetSort() sort.Sort
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
func (t *SQLTranslator) GetArgs(s statement.Statement) []interface{} {
	return t.getArgs(s.GetRoot())
}

func (t *SQLTranslator) getArgs(node nodes.Node) []interface{} {
	if node == nil {
		return []interface{}{}
	}
	switch node := node.(type) {
	case nodes.NullNode, nodes.BoolNode:
		return []interface{}{}
	case nodes.NodeWithValue:
		return []interface{}{node.Value()}
	case nodes.NodeWithChilds:
		childs := node.Childs()
		// non leaf nodes without childs are useless, ignore it
		if len(childs) == 0 {
			return []interface{}{}
		}
		args := make([]interface{}, 0)
		for _, child := range childs {
			args = append(args, t.getArgs(child)...)
		}
		return args
	}
	return nil
}

// Translate translates statement to SQL
func (t *SQLTranslator) Translate(s Statement) string {
	queryBuilder := strings.Builder{}

	root := s.GetRoot()
	if root != nil {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(t.translateNode(root))
	}
	queryBuilder.WriteString(t.translateSort(s.GetSort()))

	return queryBuilder.String()
}

func (t *SQLTranslator) translateNode(node nodes.Node) string {
	if node == nil {
		return ""
	}

	statementParentheses := false
	childsParentheses := false
	switch node.(type) {
	case nodes.AndNode, nodes.OrNode:
		statementParentheses = true
	case nodes.InNode, nodes.NotNode, nodes.NotInNode:
		childsParentheses = true
	}

	if nodeWithValue, ok := node.(nodes.NodeWithValue); ok {
		return nodeWithValue.Accept(t)
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
		// TODO: think about exclusive array node translation
		_, isArrayNode := node.(nodes.ArrayNode)
		if !isArrayNode {
			queryBuilder.WriteString(node.Accept(t))
		}
		if childsParentheses {
			queryBuilder.WriteRune(openParentheses)
		}
		translatedNode := t.translateNode(childs[0])
		if translatedNode == "" {
			return ""
		}
		queryBuilder.WriteString(translatedNode)
		if childsParentheses {
			queryBuilder.WriteRune(closeParentheses)
		}
	}
	// binary op
	if len(childs) == 2 {
		translatedNode := t.translateNode(childs[0])
		if translatedNode == "" {
			return ""
		}
		queryBuilder.WriteString(translatedNode)

		translatedNode = t.translateNode(childs[1])
		if translatedNode != "" {
			queryBuilder.WriteString(node.Accept(t))
			queryBuilder.WriteString(translatedNode)
		}
	}
	// variadic op
	if len(childs) > 2 {
		queryBuilder.WriteString(t.translateNode(childs[0]))
		for _, child := range childs[1:] {
			queryBuilder.WriteString(node.Accept(t))
			queryBuilder.WriteString(t.translateNode(child))
		}
	}
	if statementParentheses {
		queryBuilder.WriteRune(closeParentheses)
	}

	return queryBuilder.String()
}

func (t *SQLTranslator) translateSort(s sort.Sort) string {
	if s == nil {
		return ""
	}
	sortBuilder := strings.Builder{}

	sortBuilder.WriteString(" ORDER BY ")
	sortBuilder.WriteString(strings.Join(s.By(), ","))
	if s.Direction() == sort.ASCDirection {
		sortBuilder.WriteString(" ASC")
	} else {
		sortBuilder.WriteString(" DESC")
	}

	return sortBuilder.String()
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

// TranslateNotInNode translates NotIn node to sql
func (t *SQLTranslator) TranslateNotInNode(_ nodes.NotInNode) string {
	return " NOT IN "
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
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return strconv.FormatInt(int64(n), 10)
}

// TranslateFloatNode translates float node to sql
func (t *SQLTranslator) TranslateFloatNode(n nodes.FloatNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return strconv.FormatFloat(float64(n), 'f', -1, 64)
}

// TranslateUintNode translates uint node to SQL
func (t *SQLTranslator) TranslateUintNode(n nodes.UintNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return strconv.FormatUint(uint64(n), 10)
}

// TranslateBoolNode translates bool node to SQL
func (t *SQLTranslator) TranslateBoolNode(n nodes.BoolNode) string {
	if n {
		return "TRUE"
	}
	return "FALSE"
}

// TranslateLikeNode translates Like node to SQL
func (t *SQLTranslator) TranslateLikeNode(n nodes.LikeNode) string {
	return " LIKE "
}

// TranslateILikeNode translates ILike node to SQL
func (t *SQLTranslator) TranslateILikeNode(n nodes.ILikeNode) string {
	return " ILIKE "
}
