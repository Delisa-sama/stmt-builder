package limit

// Limit represents abstract limit
type Limit interface {
	Get() uint
}

// limit represents the limit of a database query
type limit struct {
	value uint
}

// Get returns the limit to a query
func (l limit) Get() uint {
	return l.value
}

// NewLimit returns a new Limit object
func NewLimit(l uint) Limit {
	return &limit{value: l}
}
