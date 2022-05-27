package main

import (
	"fmt"
	"time"

	"github.com/Delisa-sama/stmt-builder"
	"github.com/Delisa-sama/stmt-builder/operators"
	"github.com/Delisa-sama/stmt-builder/placeholders"
	"github.com/Delisa-sama/stmt-builder/translators"
)

func exampleStmt() {
	// ((id = 10 AND status = 'active') OR deleted_at IS NOT NULL)
	s := query.NewStatement("id", operators.EQOperator{}, 10).
		And(query.NewStatement("status", operators.EQOperator{}, "active")).
		Or(query.NewStatement("deleted_at", operators.NeOperator{}, nil))

	translator := translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	fmt.Println(translator.Translate(s))

	// ((id = 10 AND (status = 'active' OR deleted_at IS NOT NULL)) OR status IN ('status1','status2'))
	s = query.NewStatement("id", operators.EQOperator{}, 10).
		And(
			query.NewStatement("status", operators.EQOperator{}, "active").
				Or(query.NewStatement("deleted_at", operators.NeOperator{}, nil)),
		).Or(query.NewStatement("status", operators.InOperator{}, []string{"status1", "status2"}))
	translator = translators.NewSQLTranslator()
	fmt.Println(translator.Translate(s))
	translator = translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	fmt.Println(translator.Translate(s))

	// ((id = 10 AND created_at > '2022-05-27T13:49:32+03:00') OR !(status IN ('status1','status2')))
	s = query.NewStatement("id", operators.EQOperator{}, 10).
		And(query.NewStatement("created_at", operators.GTOperator{}, time.Now())).
		Or(query.Not(query.NewStatement("status", operators.InOperator{}, []string{"status1", "status2"})))
	translator = translators.NewSQLTranslator()
	fmt.Println(translator.Translate(s))
	translator = translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	fmt.Println(translator.Translate(s))
	fmt.Println(translator.GetArgs(s))

	//
	s = query.NewEmptyStatement()
	translator = translators.NewSQLTranslator()
	fmt.Println(translator.Translate(s))

	s = query.NewEmptyStatement().And(query.NewStatement("id", operators.EQOperator{}, 10))
	translator = translators.NewSQLTranslator()
	// id = 10
	fmt.Println(translator.Translate(s))
}

func main() {
	exampleStmt()
}
