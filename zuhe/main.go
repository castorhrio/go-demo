package main

type A struct {
	a int
}

type B struct {
	A
}

type C struct {
	*A
}

func (a A) Get() int {
	return a.a
}

func (a *A) Set(i int) {
	a.a = i
}

func main() {
	x := A{a: 1}
	y := B{
		A: x,
	}
	z := C{
		A: &x,
	}

	println(y.Get())

	y.Set(4)
	println((y.Get()))

	// B.Set(y, 39)
	(*B).Set(&y, 7)
	println(y.Get())

	C.Set(z, 10)
	println(z.Get())

	(*C).Set(&z, 34)
	println(z.Get())
}
