package main

import (
	"fmt"
	"reflect"
)

func swap[T any](a, b T) {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	if va.Kind() != reflect.Pointer {
		panic("must be pointer")
	}

	vaElem := va.Elem()
	vbElem := vb.Elem()
	tmp := vaElem.Interface()
	vaElem.Set(vbElem)
	vbElem.Set(reflect.ValueOf(tmp))
}

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)

	swap(&a, &b)

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)
}
