package main

import (
	"fmt"
	"time"

	"github.com/Delisa-sama/stmt-builder/operators"
	"github.com/Delisa-sama/stmt-builder/placeholders"
	"github.com/Delisa-sama/stmt-builder/sort"
	"github.com/Delisa-sama/stmt-builder/statement"
	"github.com/Delisa-sama/stmt-builder/translators"
	"github.com/Delisa-sama/stmt-builder/values"
)

func exampleStmt() {
	// ((id = 10 AND status = 'active') OR deleted_at IS NOT NULL)
	s := statement.New("id", operators.EQ(values.Bool(true))).
		And(statement.New("status", operators.EQ(values.String("active")))).
		Or(statement.New("deleted_at", operators.NE(values.Null())))

	translator := translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	fmt.Println(translator.Translate(s))
	fmt.Println(translator.GetArgs(s))

	// ((id = 10 AND (status = 'active' OR deleted_at IS NOT NULL)) OR status IN ('status1','status2'))
	s = statement.New("id", operators.EQ(values.Int(10))).
		And(
			statement.New("status", operators.EQ(values.String("active"))).
				Or(statement.New("deleted_at", operators.NE(values.Null()))),
		).Or(
		statement.New("status", operators.In(values.Strings("status1"))),
	)
	translator = translators.NewSQLTranslator()
	fmt.Println(translator.Translate(s))
	translator = translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	fmt.Println(translator.Translate(s))

	// ((id = 10 AND created_at > '2022-05-27T13:49:32+03:00') OR !(status IN ('status1','status2')))
	s = statement.New("id", operators.EQ(values.Int64(10))).
		And(
			statement.New("created_at", operators.GT(values.Time(time.Now()))).Or(
				statement.Not(statement.New("status", operators.In(values.Strings("status1", "status2"))))),
		)
	translator = translators.NewSQLTranslator()
	fmt.Println(translator.Translate(s))
	translator = translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	fmt.Println(translator.Translate(s))
	fmt.Println(translator.GetArgs(s))

	//
	s = statement.Empty()
	translator = translators.NewSQLTranslator()
	fmt.Println(translator.Translate(s))

	s = statement.Empty().And(statement.New("id", operators.EQ(values.Int(10))))
	translator = translators.NewSQLTranslator()
	// id = 10
	fmt.Println(translator.Translate(s))

	// NOT IN
	// type NOT IN ("archive", "block")
	s = statement.New("type", operators.NotIn(values.Strings("archive", "block"))).
		Sort(sort.By("id"), sort.DESCDirection)

	translator = translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	// type NOT IN ($1,$2)
	fmt.Println(translator.Translate(s))
	// [archive block]
	fmt.Println(translator.GetArgs(s))

	// LIMIT AND OFFSET
	s = statement.New("name", operators.EQ(values.String("Example"))).Limit(1).Offset(5)
	// WHERE name = $3 LIMIT 1 OFFSET 5
	fmt.Println(translator.Translate(s))
}

func main() {
	exampleStmt()
}
