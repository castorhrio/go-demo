package main

import (
	"fmt"
)

func chvalue(a int) int {
	a = a + 1
	return a
}

func chpointer(a *int) {
	*a = *a + 1
	return
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

type op func(int, int) int

func do(f op, a, b int) int {
	return f(a, b)
}

func main() {
	a := 10
	chvalue(a)
	fmt.Println(a)

	chpointer(&a)
	fmt.Println(a)

	fmt.Printf("%T\n", chpointer)

	v := do(add, 1, 4)
	fmt.Println(v)

	defer func() {
		fmt.Println("this is first")
	}()

	defer func() {
		fmt.Println("this is second")
	}()

	defer func(i int) {
		fmt.Println("defer i=", i)
	}(a)

	a++
	fmt.Println("result a is ", a)
}
