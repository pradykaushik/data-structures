package util

// Value stored in a data structure.
type Value interface {
	// Get the contained value.
	Get() interface{}
}
