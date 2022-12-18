package main

import (
	"fmt"

	"github.com/septemhill/enum"
)

type Result enum.SumType

type Ok = enum.SumTypeOne[int]
type Err = enum.SumTypeOne[string]

// Currently, we cannot do this.
// var OkResult = enum.NewSumTypeOne[T]
// var ErrResult = enum.NewSumTypeOne[T]

func ReturnResultWithInt() Result {
	return enum.NewSumTypeOne(123)
}

func ReturnResultWithErr() Result {
	return enum.NewSumTypeOne("error message")
}

func ResultCheck(r Result) {
	// We assume that the value type in `Ok` would be `int`.
	// And the error would be `string`.

	// Note: if make `Ok` and `Err` with the same type,
	// then we would get wrong behavior here.
	if v, ok := r.(Ok); ok {
		fmt.Println(v.Value())
	}
	if v, ok := r.(Err); ok {
		fmt.Println(v.Value())
	}
}

func main() {
	a := ReturnResultWithInt()
	ResultCheck(a)

	d := ReturnResultWithErr()
	ResultCheck(d)
}
