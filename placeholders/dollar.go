package placeholders

import (
	"strconv"
	"strings"
)

// DollarPlaceholder represents SQL numbered placeholder with dollar
type DollarPlaceholder struct {
	count int
}

// NewDollarPlaceholder returns new DollarPlaceholder
func NewDollarPlaceholder() *DollarPlaceholder {
	return &DollarPlaceholder{count: 1}
}

// Next returns dollar placeholder, starts with '$1'
func (p *DollarPlaceholder) Next() string {
	placeholderBuilder := strings.Builder{}
	placeholderBuilder.WriteRune('$')
	placeholderBuilder.WriteString(strconv.Itoa(p.count))
	p.count++
	return placeholderBuilder.String()
}
