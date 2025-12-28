package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func TrimAllStrings(a any) {
	// 記錄訪問過的目標
	visited := make(map[uintptr]struct{})

	var trim func(v reflect.Value)
	trim = func(va reflect.Value) {
		if !va.IsValid() {
			return
		}
		switch va.Kind() {
		case reflect.Pointer:
			ptr := va.Pointer()
			if _, ok := visited[ptr]; ok {
				return
			}
			visited[ptr] = struct{}{}
			trim(va.Elem())
		case reflect.Interface:
			trim(va.Elem())
		case reflect.Struct:
			for i := 0; i < va.NumField(); i++ {
				trim(va.Field(i))
			}
		case reflect.Slice, reflect.Array:
			for i := range va.Len() {
				trim(va.Index(i))
			}
		case reflect.Map:
			for _, key := range va.MapKeys() {
				trim(va.MapIndex(key))
			}
		case reflect.String:
			if va.CanSet() {
				va.SetString(strings.TrimSpace(va.String()))
			}
		}

	}

	trim(reflect.ValueOf(a))

}

func main() {
	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(&a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(&a)

	fmt.Println(a.Next.Next.Name == "name")
}
