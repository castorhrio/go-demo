package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{Id: 1, Name: "ayin", Age: 30}

	va := reflect.ValueOf(u)
	vb := reflect.ValueOf(&u)

	fmt.Println(va.CanSet(), va.FieldByName("Name").CanSet())
	fmt.Println(vb.CanSet(), vb.Elem().FieldByName("Name").CanSet())

	fmt.Printf("%v\n", vb)

	name := "zjc"
	vc := reflect.ValueOf(name)
	vb.Elem().FieldByName("Name").Set(vc)
	fmt.Printf("%v\n", vb)
}
