package main

import "fmt"

type Inter interface {
	Ping()
	Pong()
}

type Anter interface {
	Inter
	String()
}

type Student struct {
	Name string
}

func (Student) Ping() {
	println("ping")
}

func (*Student) Pong() {
	println("pong")
}

func main() {
	st := &Student{"ayin"}
	var i interface{} = st
	o := i.(Inter)
	o.Ping()
	o.Pong()

	p := i.(*Student)
	fmt.Printf("%s\n", p.Name)

	if a, ok := i.(Inter); ok {
		a.Ping()
	}

	if b, ok := i.(Anter); ok {
		b.String()
	}

	if c, ok := i.(*Student); ok {
		println(c.Name)
	}
}
