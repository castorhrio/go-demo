package main

import "fmt"

func fa(a int) func(i int) int {
	return func(i int) int {
		println(&a, a)
		a = a + 1
		return a
	}
}

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(x int) {
	t.a = x
}

func (t *T) Print() {
	fmt.Printf("%p, %v, %d\n", t, t, t.a)
}

func main() {
	f := fa(1)
	g := fa(1)

	t := &T{}
	b := T{a: 1}
	c := (T).Get(b)
	println(c)

	t.Set(5)
	println("this is type :", t.Get())
	t.Print()
	fmt.Printf("t: %v\n", t)

	println(f(1))
	println(f(1))

	println(g(1))
	println(g(1))
}
