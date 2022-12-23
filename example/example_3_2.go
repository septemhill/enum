package main

import (
	"fmt"

	"github.com/septemhill/enum"
)

func ReturnSome() enum.Option {
	return enum.SomeOption("Septem")
}

func ReturnNone() enum.Option {
	return enum.NoneOption()
}

func OptionCheck(r enum.Option) {
	if v, ok := r.(enum.Some[string]); ok {
		fmt.Println(v.Value())
	}
	if v, ok := r.(enum.None); ok {
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
