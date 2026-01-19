package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("main..")

	// int
	var variableInt8 int8 = 10
	fmt.Println("variableInt8", variableInt8, math.MaxInt8, math.MinInt8)
	var variableInt16 int16 = 10
	fmt.Println("variableInt16", variableInt16, math.MaxInt16, math.MinInt16)
	var variableInt32 int32 = 10
	fmt.Println("variableInt32", variableInt32, math.MaxInt32, math.MinInt32)
	var variableInt int = 10
	fmt.Println("variableInt", variableInt, math.MaxInt, math.MinInt)
	var variableInt64 int64 = 10
	fmt.Println("variableInt64", variableInt64, math.MaxInt64, math.MinInt64)

	// uint
	var variableUint uint = 10
	fmt.Println("variableUint", variableUint)

	var variableUint8 uint8 = 10
	fmt.Println("variableUint8", variableUint8)

	var variableUint16 uint16 = 10
	fmt.Println("variableUint16", variableUint16)

	var variableUint32 uint32 = 10
	fmt.Println("variableUint32", variableUint32)

	var variableUint64 uint64 = 10
	fmt.Println("variableUint64", variableUint64)

	// float
	var variableFloat32 float32 = 10
	fmt.Println("variableFloat32", variableFloat32)
	var variableFloat64 float64 = 10
	fmt.Println("variableFloat64", variableFloat64)

	// complex
	var variableComplex64 complex64 = 10
	fmt.Println("variableComplex64", variableComplex64)
	var variableComplex128 complex128 = 10
	fmt.Println("variableComplex128", variableComplex128)

	//bool
	var variableBool bool = false
	fmt.Println("variableBool", variableBool)

	// unicode
	var variableRune rune = 10
	fmt.Println("variableRune", variableRune)

	var variableRuneA rune = 'A'
	fmt.Printf("variableRuneA %c \n", variableRuneA)

	// string
	var variableString string = "hello world"
	fmt.Println("variableString", variableString)

}
