package main

import (
	"fmt"

	"github.com/septemhill/enum"
)

func ReturnResultWithInt() enum.Result {
	return enum.OkResult(123)
}

func ReturnResultWithErr() enum.Result {
	return enum.ErrResult("error message")
}

func ResultCheck(r enum.Result) {
	// We assume that the value type in `Ok` would be `int`.
	// And the error would be `string`.

	// Note: if make `Ok` and `Err` with the same type,
	// then we would get wrong behavior here.
	if v, ok := r.(enum.Ok[int]); ok {
		fmt.Println(v.Value())
	}
	if v, ok := r.(enum.Err[string]); ok {
		fmt.Println(v.Value())
	}
}

func main() {
	a := ReturnResultWithInt()
	ResultCheck(a)

	d := ReturnResultWithErr()
	ResultCheck(d)
}
