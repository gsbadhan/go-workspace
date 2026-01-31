package main

import "fmt"

func main() {
	fmt.Println("main..")
	fmt.Println("add(a int, b int)=", add(11, 22))
	fmt.Println("callbackType1(add func(int, int) int, a int, b int) int=", callbackType1(add, 11, 22))
	fmt.Println("callbackType2(add func(int, int) int) int=", callbackType2(add))
	calc := calculatorStruct{add: func(i1, i2 int) int { return i1 + i2 }, subtract: func(i1, i2 int) int { return i1 - i2 }, multiply: func(i1, i2 int) int { return i1 * i2 }}
	fmt.Println("calculatorStruct=", calc.add(11, 22), calc.subtract(11, 22), calc.multiply(11, 22))
	middleware1(pre, next, post)()
	middleware2(pre, next, post)
}

func add(a int, b int) int {
	return a + b
}

func callbackType1(add func(int, int) int, a int, b int) int {
	return add(a, b)
}

func callbackType2(add func(int, int) int) int {
	var a int = 11
	var b int = 22
	return add(a, b)
}

type calculatorStruct struct {
	add      func(int, int) int
	subtract func(int, int) int
	multiply func(int, int) int
}

func middleware1(pre func(), next func(), post func()) func() {
	return func() {
		pre()
		next()
		post()
	}
}

func middleware2(pre func(), next func(), post func()) {
	pre()
	next()
	post()
}

func pre()  { fmt.Print("pre ->") }
func next() { fmt.Print("next ->") }
func post() { fmt.Print("post\n") }
