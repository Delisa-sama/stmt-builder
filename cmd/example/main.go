package main

import (
	"fmt"

	"github.com/Delisa-sama/stmt-builder"
	"github.com/Delisa-sama/stmt-builder/operators"
	"github.com/Delisa-sama/stmt-builder/placeholders"
)

func exampleStmt() {
	// ((id = 10 AND status = 'active') OR deleted_at IS NOT NULL)
	s := query.NewStatement("id", operators.EQOperator{}, 10).
		And(query.NewStatement("status", operators.EQOperator{}, "active")).
		Or(query.NewStatement("deleted_at", operators.NeOperator{}, nil))
	fmt.Println(s.ToSQL(nil))
	fmt.Println(s.ToSQL(placeholders.NewDollarPlaceholder()))

	// ((id = 10 AND (status = 'active' OR deleted_at IS NOT NULL)) OR status = 'registered')
	s = query.NewStatement("id", operators.EQOperator{}, 10).
		And(
			query.NewStatement("status", operators.EQOperator{}, "active").
				Or(query.NewStatement("deleted_at", operators.NeOperator{}, nil)),
		).Or(query.NewStatement("status", operators.EQOperator{}, "registered"))
	fmt.Println(s.ToSQL(nil))
	fmt.Println(s.ToSQL(placeholders.NewQuestionMarkPlaceholder()))
}

func main() {
	exampleStmt()
}
