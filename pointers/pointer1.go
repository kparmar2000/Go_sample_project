package main

import (
	"fmt"
	"reflect"
)

func main() {
	//x := 10
	//fmt.Println(x)
	A := 'A'
	ptr := new(int)
	changenumber(ptr)
	fmt.Println(A)
	fmt.Println(reflect.TypeOf(A))
	a := 123
	var p *int = &a
	fmt.Println(p)

}
func changenumber(p *int) {
	*p = 10
}