package tl

type Optional[T any] struct {
	value    T
	hasValue bool
}

func MakeOptional[T any](v T) (opt Optional[T]) {
	opt.Set(v)
	return opt
}

func NewOptional[T any](v *T) (opt Optional[T]) {
	if v != nil {
		opt.Set(*v)
	}

	return opt
}

func None[T any]() (opt Optional[T]) {
	return opt
}

func (opt Optional[T]) From(v T) Optional[T] {
	opt.Set(v)
	return opt
}

func (opt Optional[T]) Ptr() *T {
	if opt.hasValue {
		return &opt.value
	}

	return nil
}

func (opt Optional[T]) Get() T {
	return opt.value
}

func (opt Optional[T]) Or(v T) Optional[T] {
	if !opt.hasValue {
		opt.Set(v)
	}

	return opt
}

func (opt Optional[T]) HasValue() bool {
	return opt.hasValue
}

func (opt *Optional[T]) Set(v T) {
	opt.value = v
	opt.hasValue = true
}

func (opt *Optional[T]) Reset() {
	opt.hasValue = false
}

type OptionalPtr[T any] struct {
	value *T
}

func MakeOptionalPtr[T any](v *T) (opt OptionalPtr[T]) {
	if v != nil {
		opt.Set(v)
	}

	return opt
}

func (opt OptionalPtr[T]) Ptr() *T {
	if opt.value != nil {
		return opt.value
	}

	return nil
}

func (opt OptionalPtr[T]) Get() *T {
	return opt.value
}

func (opt *OptionalPtr[T]) GetValue() T {
	return *opt.value
}

func (opt OptionalPtr[T]) Or(v *T) OptionalPtr[T] {
	if !opt.HasValue() {
		opt.Set(v)
	}

	return opt
}

func (opt OptionalPtr[T]) HasValue() bool {
	return opt.value != nil
}

func (opt *OptionalPtr[T]) Set(v *T) {
	opt.value = v
}

func (opt *OptionalPtr[T]) Reset() {
	opt.value = nil
}
