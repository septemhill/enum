package enum

type Result SumType

type Ok[T any] interface {
	Value() T
}

type Err[T any] interface {
	Value() T
}

type OkImpl[T any] struct {
	value T
}

func (st *OkImpl[T]) Value() T {
	return st.value
}

func OkResult[T any](value T) *OkImpl[T] {
	return &OkImpl[T]{
		value: value,
	}
}

type ErrImpl[T any] struct {
	value T
}

func (st *ErrImpl[T]) Value() T {
	return st.value
}

func ErrResult[T any](value T) *ErrImpl[T] {
	return &ErrImpl[T]{
		value: value,
	}
}
