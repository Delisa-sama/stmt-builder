package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/stmt-builder/operators"
	"github.com/Delisa-sama/stmt-builder/placeholders"
	"github.com/Delisa-sama/stmt-builder/sort"
	"github.com/Delisa-sama/stmt-builder/statement"
	"github.com/Delisa-sama/stmt-builder/values"
)

func TestSQLTranslator_Translate(t *testing.T) {
	tests := []struct {
		name        string
		placeholder Placeholder
		s           statement.Statement
		want        string
	}{
		{
			name:        "single statement expression without brackets",
			placeholder: nil,
			s:           statement.New("id", operators.EQ(values.Int(10))),
			want:        " WHERE id = 10",
		},
		{
			name:        "empty statement",
			placeholder: nil,
			s:           statement.Empty(),
			want:        "",
		},
		{
			name:        "complex expression without placeholder",
			placeholder: nil,
			s: statement.New("id", operators.EQ(values.Int(10))).
				And(statement.New("status", operators.In(values.Strings("active", "blocked")))),
			want: " WHERE (id = 10 AND status IN ('active','blocked'))",
		},
		{
			name:        "complex expression with dollar placeholder",
			placeholder: placeholders.NewDollarPlaceholder(),
			s: statement.New("id", operators.EQ(values.Int(10))).
				And(statement.New("status", operators.In(values.Strings("active", "blocked")))),
			want: " WHERE (id = $1 AND status IN ($2,$3))",
		},
		{
			name:        "complex expression with question mark placeholder",
			placeholder: placeholders.NewQuestionMarkPlaceholder(),
			s: statement.New("id", operators.EQ(values.Int(10))).
				And(statement.New("status", operators.In(values.Strings("active", "blocked")))),
			want: " WHERE (id = ? AND status IN (?,?))",
		},
		{
			name:        "more complex expression",
			placeholder: nil,
			s: statement.Not(
				statement.New("id", operators.EQ(values.Int(10))).
					And(statement.New("status", operators.NE(values.String("active")))).
					Or(statement.New("deleted_at", operators.EQ(values.Null()))),
			).And(
				statement.New("weight", operators.LT(values.Float(25.123))).
					Or(statement.New("weight", operators.GE(values.Float(12.0)))),
			),
			want: " WHERE (!(((id = 10 AND status <> 'active') OR deleted_at IS NULL)) AND (weight < 25.123 OR weight >= 12))",
		},
		{
			name:        "single statement with sort",
			placeholder: nil,
			s:           statement.New("id", operators.NE(values.Int(10))).Sort([]string{"id"}, sort.DESCDirection),
			want:        " WHERE id <> 10 ORDER BY id DESC",
		},
		{
			name:        "complex expression with multiple column sort",
			placeholder: placeholders.NewQuestionMarkPlaceholder(),
			s: statement.New("id", operators.EQ(values.Int(10))).
				And(statement.New("status", operators.In(values.Strings("active", "blocked")))).
				Sort(sort.By("id", "status"), sort.DESCDirection),
			want: " WHERE (id = ? AND status IN (?,?)) ORDER BY id,status DESC",
		},
		{
			name:        "single statement LIKE expression",
			placeholder: nil,
			s:           statement.New("name", operators.LIKE("%abc%")),
			want:        " WHERE name LIKE '%abc%'",
		},
		{
			name:        "single statement LIKE expression",
			placeholder: nil,
			s:           statement.New("name", operators.LIKE("%abc%")),
			want:        " WHERE name LIKE '%abc%'",
		},
		{
			name:        "single statement ILIKE expression",
			placeholder: nil,
			s:           statement.New("name", operators.ILIKE("%abc%")).Limit(1).Offset(5),
			want:        " WHERE name ILIKE '%abc%' LIMIT 1 OFFSET 5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translator := NewTranslator(WithPlaceholder(tt.placeholder))
			got := translator.Translate(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSQLTranslator_GetArgs(t *testing.T) {
	tests := []struct {
		name        string
		placeholder Placeholder
		s           statement.Statement
		want        []interface{}
	}{
		{
			name: "single statement expression",
			s:    statement.New("id", operators.EQ(values.Int(10))),
			want: []interface{}{int64(10)},
		},
		{
			name: "empty statement",
			s:    statement.Empty(),
			want: []interface{}{},
		},
		{
			name: "more complex expression",
			s: statement.Not(
				statement.New("id", operators.EQ(values.Int(10))).
					And(statement.New("status", operators.NE(values.String("active")))).
					Or(statement.New("deleted_at", operators.EQ(values.Null()))),
			).And(
				statement.New("weight", operators.LT(values.Float(25.123))).
					Or(statement.New("weight", operators.GE(values.Float(12.0)))),
			),
			want: []interface{}{
				int64(10),
				"active",
				25.123,
				12.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translator := &Translator{
				placeholder: tt.placeholder,
			}
			got := translator.GetArgs(tt.s)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
