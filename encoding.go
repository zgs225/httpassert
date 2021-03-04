package httpassert

// Encoding encode/decode object
type Encoding interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}
