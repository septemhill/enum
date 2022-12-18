package enum

type SumType interface{}

type SumTypeZero interface {
	Value()
}

type SumTypeOne[A any] interface {
	Value() A
}

type SumTypeTwo[A, B any] interface {
	Value() (A, B)
}

type SumTypeThree[A, B, C any] interface {
	Value() (A, B, C)
}

type SumTypeFour[A, B, C, D any] interface {
	Value() (A, B, C, D)
}

type SumTypeFive[A, B, C, D, E any] interface {
	Value() (A, B, C, D, E)
}
