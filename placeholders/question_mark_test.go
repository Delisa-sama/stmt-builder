package placeholders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestionMarkPlaceholder_Next(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "next ph is '?'",
			want: "?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &QuestionMarkPlaceholder{}
			assert.Equal(t, tt.want, p.Next())
		})
	}
}
