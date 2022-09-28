package tl

import (
	"bytes"
	"io"
	"unsafe"
)

// Bytes wraps the behavior of Vec as Vec[byte] adding new methods on top.
type Bytes Vec[byte]

var (
	_ io.WriterTo   = Bytes{}
	_ io.ReaderFrom = Bytes{}
)

// NewBytes returns a Bytes instance with size `size` and capacity `capacity`.
func NewBytes(size, capacity int) Bytes {
	return BytesFrom(make([]byte, size, capacity))
}

// BytesFrom creates a Bytes instance from `bts` bytes.
func BytesFrom(bts []byte) Bytes {
	return Bytes(bts)
}

// Slice returns a sliced Bytes using indexes starting from `low` and ending in `max`.
func (b Bytes) Slice(low, max int) Bytes {
	if max <= 0 {
		return b[low:]
	}

	return b[low:max]
}

// CopyFrom copies the bytes from `other` to `b`.
func (b *Bytes) CopyFrom(other Bytes) {
	b.Reserve(other.Len())
	copy(*b, other)
}

// WriteTo writes the bytes to an io.Writer.
//
// Implements the io.WriterTo interface.
func (b Bytes) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b)

	return int64(n), err
}

// ReadFrom reads bytes from an io.Reader into `b`.
//
// This function does NOT append bytes, it uses the existing buffer.
//
// Implements the io.ReaderFrom interface.
func (b Bytes) ReadFrom(r io.Reader) (int64, error) {
	n, err := r.Read(b)

	return int64(n), err
}

// LimitReadFrom is like ReadFrom but it is limited to `n` bytes.
func (b *Bytes) LimitReadFrom(r io.Reader, n int) (int64, error) {
	b.Reserve(n)
	return b.Slice(0, n).ReadFrom(r)
}

// LimitWriteTo writes a limited amount of `n` bytes to an io.Writer.
func (b Bytes) LimitWriteTo(w io.Writer, n int) (int64, error) {
	return b.Slice(0, Min(b.Len(), n)).WriteTo(w)
}

// Len returns the length of `b`.
//
// This function is equivalent to `len(b))`.
func (b Bytes) Len() int {
	return len(b)
}

// Cap returns the capacity of `b`.
//
// This function is equivalent to `cap(b))`.
func (b Bytes) Cap() int {
	return cap(b)
}

// String converts `b` to a string.
//
// This function produces an allocation. To avoid the allocation use UnsafeString.
func (b Bytes) String() string {
	return string(b)
}

// UnsafeString converts `b` to a string without producing allocations.
func (b Bytes) UnsafeString() string {
	return *(*string)(unsafe.Pointer(&b))
}

// IndexByte returns the index of `c` in `b`.
//
// This function is equivalent to bytes.IndexByte(b, c).
func (b Bytes) IndexByte(c byte) int {
	return bytes.IndexByte(b, c)
}

// Index returns the index of `other` in `b`.
//
// This function is equivalent to bytes.Index(b, other).
func (b Bytes) Index(other Bytes) int {
	return bytes.Index(b, other)
}

// Resize increases or decreases the size of Bytes.
func (b *Bytes) Resize(n int) {
	vc := (*Vec[byte])(b)
	vc.Resize(n)
}

// Reserve increases the size of Bytes if needed.
//
// The call doesn't produce any allocation if `n` < cap(b).
func (b *Bytes) Reserve(n int) {
	vc := (*Vec[byte])(b)
	vc.Reserve(n)
}

// Append appends bytes at the end of Bytes.
func (b *Bytes) Append(others ...byte) {
	vc := (*Vec[byte])(b)
	vc.Append(others...)
}

// Push pushes the `bts` to the beginning of Bytes.
func (b *Bytes) Push(bts ...byte) {
	vc := (*Vec[byte])(b)
	vc.Push(bts...)
}
