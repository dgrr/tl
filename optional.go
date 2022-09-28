package tl

// Optional defines a value that can be optional. So, if might be defined or not.
type Optional[T any] struct {
	value    T
	hasValue bool
}

// MakeOptional returns an Optional with a value `v`.
func MakeOptional[T any](v T) (opt Optional[T]) {
	opt.Set(v)
	return opt
}

// NewOptional returns an Optional with a value `v` if `v` is not nil.
func NewOptional[T any](v *T) (opt Optional[T]) {
	if v != nil {
		opt.Set(*v)
	}

	return opt
}

// None returns an empty Optional.
func None[T any]() (opt Optional[T]) {
	return opt
}

// From sets the value `v` to the Optional.
func (opt Optional[T]) From(v T) Optional[T] {
	opt.Set(v)
	return opt
}

// Ptr returns a pointer to the value or nil if there's no value.
func (opt Optional[T]) Ptr() *T {
	if opt.hasValue {
		return &opt.value
	}

	return nil
}

// Get returns the value. The function doesn't check whether the value is set or not.
func (opt Optional[T]) Get() T {
	return opt.value
}

// Or sets the value `v` to the Optional only if no value is already present.
func (opt Optional[T]) Or(v T) Optional[T] {
	if !opt.hasValue {
		opt.Set(v)
	}

	return opt
}

// HasValue returns true if Optional contains a valid value.
func (opt Optional[T]) HasValue() bool {
	return opt.hasValue
}

// Set sets a value `v` to the Optional.
func (opt *Optional[T]) Set(v T) {
	opt.value = v
	opt.hasValue = true
}

// Reset sets the Optional as invalid.
func (opt *Optional[T]) Reset() {
	opt.hasValue = false
}

// OptionalPtr is like Optional but instead of holding a copy of the value, it holds a pointer.
type OptionalPtr[T any] struct {
	value *T
}

// MakeOptionalPtr returns an OptionalPtr using the value `v`.
func MakeOptionalPtr[T any](v *T) (opt OptionalPtr[T]) {
	opt.Set(v)
	return opt
}

// Ptr returns a pointer to the value.
func (opt OptionalPtr[T]) Ptr() *T {
	return opt.value
}

// Get returns a pointer to the value, because the value is a pointer already.
func (opt OptionalPtr[T]) Get() *T {
	return opt.value
}

// GetValue returns the value inside the Optional's pointer. This function might panic
// if the value is not defined (aka value == nil).
func (opt *OptionalPtr[T]) GetValue() T {
	return *opt.value
}

// Or sets the value `v` to the Optional only if not value is already present.
func (opt OptionalPtr[T]) Or(v *T) OptionalPtr[T] {
	if !opt.HasValue() {
		opt.Set(v)
	}

	return opt
}

// HasValue returns true if a value is present.
func (opt OptionalPtr[T]) HasValue() bool {
	return opt.value != nil
}

// Set sets the value `v` to the OptionalPtr.
func (opt *OptionalPtr[T]) Set(v *T) {
	opt.value = v
}

// Reset resets the OptionalPtr's value.
func (opt *OptionalPtr[T]) Reset() {
	opt.value = nil
}
