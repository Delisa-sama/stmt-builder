package offset

// Offset represents abstract offset
type Offset interface {
	Get() uint
}

// offset represents the offset of a database query
type offset struct {
	value uint
}

// Get returns the offset to a query
func (l offset) Get() uint {
	return l.value
}

// NewOffset returns a new Offset object
func NewOffset(value uint) Offset {
	return &offset{value: value}
}
