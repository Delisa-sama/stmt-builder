package placeholders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDollarPlaceholder_Next(t *testing.T) {
	tests := []struct {
		name  string
		count int
		want  string
	}{
		{
			name:  "first ph is '$1'",
			count: 1,
			want:  "$1",
		},
		{
			name:  "second ph is '$2'",
			count: 2,
			want:  "$2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeholder := &DollarPlaceholder{count: tt.count}
			ph := placeholder.Next()
			assert.Equal(t, tt.want, ph)
		})
	}
}
