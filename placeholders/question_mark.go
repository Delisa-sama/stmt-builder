package placeholders

// QuestionMarkPlaceholder represents placeholder with '?' symbol
type QuestionMarkPlaceholder struct{}

// Next returns '?' placeholder
func (p *QuestionMarkPlaceholder) Next() string {
	return "?"
}

// NewQuestionMarkPlaceholder returns new QuestionMarkPlaceholder
func NewQuestionMarkPlaceholder() *QuestionMarkPlaceholder {
	return &QuestionMarkPlaceholder{}
}
