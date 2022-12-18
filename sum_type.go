package enum

type SumTypeZeroImpl struct{}

func (st *SumTypeZeroImpl) Value() {}

func NewSumTypeZero() *SumTypeZeroImpl {
	return &SumTypeZeroImpl{}
}

type SumTypeOneImpl[A any] struct {
	valueA A
}

func (st *SumTypeOneImpl[A]) Value() A {
	return st.valueA
}

func NewSumTypeOne[A any](valueA A) *SumTypeOneImpl[A] {
	return &SumTypeOneImpl[A]{
		valueA: valueA,
	}
}

type SumTypeTwoImpl[A, B any] struct {
	valueA A
	valueB B
}

func (st *SumTypeTwoImpl[A, B]) Value() (A, B) {
	return st.valueA, st.valueB
}

func NewSumTypeTwo[A, B any](
	valueA A,
	valueB B,
) *SumTypeTwoImpl[A, B] {
	return &SumTypeTwoImpl[A, B]{
		valueA: valueA,
		valueB: valueB,
	}
}

type SumTypeThreeImpl[A, B, C any] struct {
	valueA A
	valueB B
	valueC C
}

func (st *SumTypeThreeImpl[A, B, C]) Value() (A, B, C) {
	return st.valueA, st.valueB, st.valueC
}

func NewSumTypeThree[A, B, C any](
	valueA A,
	valueB B,
	valueC C,
) *SumTypeThreeImpl[A, B, C] {
	return &SumTypeThreeImpl[A, B, C]{
		valueA: valueA,
		valueB: valueB,
		valueC: valueC,
	}
}

type SumTypeFourImpl[A, B, C, D any] struct {
	valueA A
	valueB B
	valueC C
	valueD D
}

func (st *SumTypeFourImpl[A, B, C, D]) Value() (A, B, C, D) {
	return st.valueA, st.valueB, st.valueC, st.valueD
}

func NewSumTypeFour[A, B, C, D any](
	valueA A,
	valueB B,
	valueC C,
	valueD D,
) *SumTypeFourImpl[A, B, C, D] {
	return &SumTypeFourImpl[A, B, C, D]{
		valueA: valueA,
		valueB: valueB,
		valueC: valueC,
		valueD: valueD,
	}
}

type SumTypeFiveImpl[A, B, C, D, E any] struct {
	valueA A
	valueB B
	valueC C
	valueD D
	valueE E
}

func (st *SumTypeFiveImpl[A, B, C, D, E]) Value() (A, B, C, D, E) {
	return st.valueA, st.valueB, st.valueC, st.valueD, st.valueE
}

func NewSumTypeFive[A, B, C, D, E any](
	valueA A,
	valueB B,
	valueC C,
	valueD D,
	valueE E,
) *SumTypeFiveImpl[A, B, C, D, E] {
	return &SumTypeFiveImpl[A, B, C, D, E]{
		valueA: valueA,
		valueB: valueB,
		valueC: valueC,
		valueD: valueD,
		valueE: valueE,
	}
}
