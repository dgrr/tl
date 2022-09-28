package tl

import (
	"bytes"
	"io"
	"unsafe"
)

type Bytes Vec[byte]

var (
	_ io.WriterTo   = Bytes{}
	_ io.ReaderFrom = Bytes{}
)

func NewBytes(size, capacity int) Bytes {
	return BytesFrom(make([]byte, size, capacity))
}

func BytesFrom(bts []byte) Bytes {
	return Bytes(bts)
}

func (b Bytes) Slice(low, max int) Bytes {
	if max <= 0 {
		return b[low:]
	}

	return b[low:max]
}

func (b *Bytes) CopyFrom(b2 Bytes) {
	b.Reserve(b2.Len())
	copy(*b, b2)
}

func (b Bytes) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b)

	return int64(n), err
}

func (b Bytes) ReadFrom(r io.Reader) (int64, error) {
	n, err := r.Read(b)

	return int64(n), err
}

func (b *Bytes) LimitReadFrom(r io.Reader, n int) (int64, error) {
	b.Reserve(n)
	return b.Slice(0, n).ReadFrom(r)
}

// LimitWriteTo writes a limited amount from `b` to `w`.
func (b Bytes) LimitWriteTo(w io.Writer, n int) (int64, error) {
	return b.Slice(0, Min(b.Len(), n)).WriteTo(w)
}

func (b Bytes) Len() int {
	return len(b)
}

func (b Bytes) Cap() int {
	return cap(b)
}

// String returns the `Bytes`' string representation.
func (b Bytes) String() string {
	return string(b)
}

func (b Bytes) UnsafeString() string {
	return *(*string)(unsafe.Pointer(&b))
}

func (b Bytes) Index(c byte) int {
	return bytes.IndexByte(b, c)
}

func (b Bytes) Contains(c byte) bool {
	return bytes.IndexByte(b, c) != -1
}

func (b *Bytes) Resize(n int) {
	vc := (*Vec[byte])(b)
	vc.Resize(n)
}

func (b *Bytes) Reserve(n int) {
	vc := (*Vec[byte])(b)
	vc.Reserve(n)
}

func (b *Bytes) Append(bts ...byte) {
	vc := (*Vec[byte])(b)
	vc.Append(bts...)
}

func (b *Bytes) Push(bts ...byte) {
	vc := (*Vec[byte])(b)
	vc.Push(bts...)
}
