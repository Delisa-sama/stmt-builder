package placeholders

type QuestionMarkPlaceholder struct{}

func (p *QuestionMarkPlaceholder) Next() string {
	return "?"
}

func NewQuestionMarkPlaceholder() *QuestionMarkPlaceholder {
	return &QuestionMarkPlaceholder{}
}
