package gee_cache

// ByteView holds an immutable view of bytes
type ByteView struct {
	b []byte
}

// Len return view length
func (b ByteView) Len() int {
	return len(b.b)
}

// ByteSlice returns a copy of the data as a byte slice.
func (b ByteView) ByteSlice() []byte {
	return copyBytes(b.b)
}

// String returns the data as a string, making a copy if necessary.
func (b ByteView) String() string {
	return string(b.b)
}

// copyBytes return copy []byte
func copyBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
