package sort

// Sort represents abstract sort
type Sort interface {
	By() []string
	Direction() Direction
}

// Direction represents sort direction
type Direction bool

// All possible direction values
const (
	ASCDirection  Direction = true
	DESCDirection Direction = false
)

// ColumnNames represents array of column names
type ColumnNames []string

// By returns column names from strings
func By(c ...string) ColumnNames {
	return c
}

type sort struct {
	columnNames []string
	direction   Direction
}

// NewSort returns sort object
func NewSort(columnNames []string, direction Direction) Sort {
	return &sort{
		columnNames: columnNames,
		direction:   direction,
	}
}

func (s *sort) By() []string {
	return s.columnNames
}

func (s *sort) Direction() Direction {
	return s.direction
}
