package sql

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Delisa-sama/stmt-builder/limit"
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/offset"
	"github.com/Delisa-sama/stmt-builder/sort"
	"github.com/Delisa-sama/stmt-builder/statement"
	"github.com/Delisa-sama/stmt-builder/translators/sql/frame"
	generic_stack "github.com/Delisa-sama/stmt-builder/translators/stack"
)

// Placeholder represents abstract placeholder for SQL query
type Placeholder interface {
	Next() string
}

// Translator represents translator from statement to SQL
type Translator struct {
	placeholder Placeholder
}

// Statement represents abstract statement
type Statement interface {
	GetRoot() nodes.Node
	GetSort() sort.Sort
	GetLimit() limit.Limit
	GetOffset() offset.Offset
}

// TranslatorOption represents option for Translator
type TranslatorOption func(translator *Translator)

// WithPlaceholder option to define placeholder for Translator
func WithPlaceholder(placeholder Placeholder) TranslatorOption {
	return func(translator *Translator) {
		translator.placeholder = placeholder
	}
}

// NewTranslator returns new Translator
func NewTranslator(opts ...TranslatorOption) *Translator {
	t := &Translator{}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

// GetArgs returns args for SQL query from statement
func (t *Translator) GetArgs(s statement.Statement) []interface{} {
	return t.getArgs(s.GetRoot())
}

func (t *Translator) getArgs(node nodes.Node) []interface{} {
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
func (t *Translator) Translate(s Statement) string {
	queryBuilder := strings.Builder{}

	root := s.GetRoot()
	if root != nil {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(t.translateNode(root))
	}
	queryBuilder.WriteString(t.translateSort(s.GetSort()))

	queryBuilder.WriteString(t.translateLimit(s.GetLimit()))

	queryBuilder.WriteString(t.translateOffset(s.GetOffset()))

	return queryBuilder.String()
}

func (t *Translator) translateNode(root nodes.Node) string {
	start := frame.NewFrame(root, nil, 0)
	stack := generic_stack.Stack[*frame.Frame]{start}

	for !stack.IsEmpty() {
		var returnValue string
		f := stack.Pop()
		if child := f.GetNextChild(); child != nil {
			stack.Push(f)
			stack.Push(child)
		} else {
			returnValue = f.TranslateFrame(t)
		}

		f.PassToParent(returnValue)
	}

	return start.Result()
}

func (t *Translator) translateSort(s sort.Sort) string {
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

func (t *Translator) translateLimit(l limit.Limit) string {
	if l == nil {
		return ""
	}
	limitBuilder := strings.Builder{}

	limitBuilder.WriteString(" LIMIT ")
	limitBuilder.WriteString(strconv.FormatUint(uint64(l.Get()), 10))

	return limitBuilder.String()
}

func (t *Translator) translateOffset(o offset.Offset) string {
	if o == nil {
		return ""
	}
	offsetBuilder := strings.Builder{}

	offsetBuilder.WriteString(" OFFSET ")
	offsetBuilder.WriteString(strconv.FormatUint(uint64(o.Get()), 10))

	return offsetBuilder.String()
}

// TranslateAndNode translates and node to sql
func (t *Translator) TranslateAndNode(node nodes.AndNode) string {
	return " AND "
}

// TranslateArrayNode translates array node to sql
func (t *Translator) TranslateArrayNode(node nodes.ArrayNode) string {
	return ","
}

// TranslateEqNode translates EQ node to sql
func (t *Translator) TranslateEqNode(node nodes.EqNode) string {
	_, isNull := node.Right().(nodes.NullNode)
	if isNull {
		return " IS "
	}
	return " = "
}

// TranslateGeNode translates GE node to sql
func (t *Translator) TranslateGeNode(node nodes.GeNode) string {
	return " >= "
}

// TranslateGtNode translates GT node to sql
func (t *Translator) TranslateGtNode(node nodes.GtNode) string {
	return " > "
}

// TranslateInNode translates In node to sql
func (t *Translator) TranslateInNode(node nodes.InNode) string {
	return " IN "
}

// TranslateNotInNode translates NotIn node to sql
func (t *Translator) TranslateNotInNode(_ nodes.NotInNode) string {
	return " NOT IN "
}

// TranslateLeNode translates LE node to sql
func (t *Translator) TranslateLeNode(node nodes.LeNode) string {
	return " <= "
}

// TranslateLtNode translates LT node to sql
func (t *Translator) TranslateLtNode(node nodes.LtNode) string {
	return " < "
}

// TranslateNameNode translates name node to sql
func (t *Translator) TranslateNameNode(node nodes.NameNode) string {
	return node.Name()
}

// TranslateNeNode translates ne node to sql
func (t *Translator) TranslateNeNode(node nodes.NeNode) string {
	_, isNull := node.Right().(nodes.NullNode)
	if isNull {
		return " IS NOT "
	}
	return " <> "
}

// TranslateNotNode translates not node to sql
func (t *Translator) TranslateNotNode(node nodes.NotNode) string {
	return "!"
}

// TranslateNullNode translates null node to sql
func (t *Translator) TranslateNullNode(node nodes.NullNode) string {
	return "NULL"
}

// TranslateOrNode translates or node to sql
func (t *Translator) TranslateOrNode(node nodes.OrNode) string {
	return " OR "
}

// TranslateStringNode translates string node to sql
func (t *Translator) TranslateStringNode(node nodes.StringNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return fmt.Sprintf("'%s'", node)
}

// TranslateValueNode translates value node to sql
func (t *Translator) TranslateValueNode(node nodes.ValueNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return fmt.Sprintf("%v", node.Value())
}

// TranslateTimeNode translates time node to sql
func (t *Translator) TranslateTimeNode(node nodes.TimeNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return fmt.Sprintf("'%s'", node.Format(time.RFC3339))
}

// TranslateIntNode translates int node to sql
func (t *Translator) TranslateIntNode(n nodes.IntNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return strconv.FormatInt(int64(n), 10)
}

// TranslateFloatNode translates float node to sql
func (t *Translator) TranslateFloatNode(n nodes.FloatNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return strconv.FormatFloat(float64(n), 'f', -1, 64)
}

// TranslateUintNode translates uint node to SQL
func (t *Translator) TranslateUintNode(n nodes.UintNode) string {
	if t.placeholder != nil {
		return t.placeholder.Next()
	}
	return strconv.FormatUint(uint64(n), 10)
}

// TranslateBoolNode translates bool node to SQL
func (t *Translator) TranslateBoolNode(n nodes.BoolNode) string {
	if n {
		return "TRUE"
	}
	return "FALSE"
}

// TranslateLikeNode translates Like node to SQL
func (t *Translator) TranslateLikeNode(n nodes.LikeNode) string {
	return " LIKE "
}

// TranslateILikeNode translates ILike node to SQL
func (t *Translator) TranslateILikeNode(n nodes.ILikeNode) string {
	return " ILIKE "
}
