package main

import (
	"fmt"

	"github.com/codegangsta/inject"
)

type A interface{}
type B interface{}

type Student struct {
	Name    string `inject`
	Age     int    `inject`
	Sex     A      `inject`
	Address B      `inject`
}

func main() {
	s := Student{}

	inj := inject.New()
	inj.Map("ayin")
	inj.Map(23)
	inj.MapTo("man", (*A)(nil))
	inj.MapTo("yuna", (*B)(nil))

	inj.Apply(&s)

	fmt.Printf("s=%v\n", s)
}
