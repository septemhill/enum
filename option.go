package enum

type Option SumType

// Same implementation with SumTypeOne
type Some[T any] interface {
	Value() T
}

type None = SumTypeZero

type SomeImpl[T any] struct {
	value T
}

func (st *SomeImpl[T]) Value() T {
	return st.value
}

func SomeOption[T any](value T) *SomeImpl[T] {
	return &SomeImpl[T]{
		value: value,
	}
}

var NoneOption = NewSumTypeZero
