
## TODO
- [x] strongly typed nodes
- [ ] brackets normalization
- [ ] non-recursive AST traverse

## Install
```shell
go get github.com/Delisa-sama/stmt-builder@latest
```

## Usage
```go
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
	s := query.NewStatement("id", operators.EQOperator{}, 10).
		And(query.NewStatement("status", operators.EQOperator{}, "active")).
		Or(query.NewStatement("deleted_at", operators.NeOperator{}, nil))

	translator := translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	// ((id = 10 AND status = $1) OR deleted_at IS NOT NULL)
	fmt.Println(translator.Translate(s))

	s = query.NewStatement("id", operators.EQOperator{}, 10).
		And(
			query.NewStatement("status", operators.EQOperator{}, "active").
				Or(query.NewStatement("deleted_at", operators.NeOperator{}, nil)),
		).Or(query.NewStatement("status", operators.InOperator{}, []string{"status1", "status2"}))
	translator = translators.NewSQLTranslator()
	// ((id = 10 AND (status = 'active' OR deleted_at IS NOT NULL)) OR status IN ('status1','status2'))
	fmt.Println(translator.Translate(s))
	translator = translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	// ((id = 10 AND (status = $1 OR deleted_at IS NOT NULL)) OR status IN ($2,$3))
	fmt.Println(translator.Translate(s))

	s = query.NewStatement("id", operators.EQOperator{}, 10).
		And(query.NewStatement("created_at", operators.GTOperator{}, time.Now())).
		Or(query.Not(query.NewStatement("status", operators.InOperator{}, []string{"status1", "status2"})))
	translator = translators.NewSQLTranslator()
	// ((id = 10 AND created_at > '2022-05-27T16:55:04+03:00') OR !(status IN ('status1','status2')))
	fmt.Println(translator.Translate(s))
	translator = translators.NewSQLTranslator(
		translators.WithPlaceholder(placeholders.NewDollarPlaceholder()),
	)
	// ((id = 10 AND created_at > $1) OR !(status IN ($2,$3)))
	fmt.Println(translator.Translate(s))
	// [10 2022-05-27 16:55:04.780222 +0300 MSK m=+0.000167668 status1 status2]
	fmt.Println(translator.GetArgs(s))
}

func main() {
	exampleStmt()
}
```