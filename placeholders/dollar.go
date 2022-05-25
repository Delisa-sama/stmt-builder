package placeholders

import (
	"strconv"
	"strings"
)

type DollarPlaceholder struct {
	count int
}

func NewDollarPlaceholder() *DollarPlaceholder {
	return &DollarPlaceholder{count: 1}
}

func (p *DollarPlaceholder) Next() string {
	placeholderBuilder := strings.Builder{}
	placeholderBuilder.WriteRune('$')
	placeholderBuilder.WriteString(strconv.Itoa(p.count))
	p.count++
	return placeholderBuilder.String()
}
