package main

import (
	"fmt"

	"github.com/septemhill/enum"
)

type Color enum.SumType

type Red = enum.SumTypeZero
type Green = enum.SumTypeTwo[int, int]
type Yellow = enum.SumTypeThree[string, int, float64]
type Blue = enum.SumTypeThree[int, string, []byte]

var RedColor = enum.NewSumTypeZero
var GreenColor = enum.NewSumTypeTwo[int, int]
var YellowColor = enum.NewSumTypeThree[string, int, float64]
var BlueColor = enum.NewSumTypeThree[int, string, []byte]

func ReturnRed() Color {
	return RedColor()
}

func ReturnGreen() Color {
	return GreenColor(321, 444)
}

func ReturnBlue() Color {
	return BlueColor(321, "[Blue]", []byte("[Blue]"))
}

func ReturnYellow() Color {
	return YellowColor("[Yellow]", 222, 44.33)
}

func EnumCheck(e Color) {
	if v, ok := e.(Red); ok {
		// Value here is meaningless
		v.Value()
	}
	if v, ok := e.(Green); ok {
		fmt.Println(v.Value())
	}
	if v, ok := e.(Yellow); ok {
		fmt.Println(v.Value())
	}
	if v, ok := e.(Blue); ok {
		fmt.Println(v.Value())
	}
}

func main() {
	g := ReturnGreen()
	EnumCheck(g)

	b := ReturnBlue()
	EnumCheck(b)

	y := ReturnYellow()
	EnumCheck(y)

	r := ReturnRed()
	EnumCheck(r)
}
