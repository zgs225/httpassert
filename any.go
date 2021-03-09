package httpassert

// Any compare to any value returns true
type Any struct{}

// Equal compare to i
func (Any) Equal(i interface{}) bool {
	return true
}
