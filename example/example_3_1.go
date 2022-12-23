package main

import (
	"fmt"

	"github.com/septemhill/enum"
)

type Option enum.SumType

type Some = enum.SumTypeOne[string]
type None = enum.SumTypeZero

// Currently, we cannot do this.
// var SomeOption = enum.NewSumTypeOne[T]
var NoneOption = enum.NewSumTypeZero

func ReturnSome() Option {
	return enum.NewSumTypeOne("Septem")
}

func ReturnNone() Option {
	return NoneOption()
}

func OptionCheck(r Option) {
	if v, ok := r.(Some); ok {
		fmt.Println(v.Value())
	}
	if v, ok := r.(None); ok {
		// None is enum.SumTypeZero which Value() without returns.
		v.Value()
	}
}

func main() {
	a := ReturnSome()
	OptionCheck(a)

	d := ReturnNone()
	OptionCheck(d)
}
