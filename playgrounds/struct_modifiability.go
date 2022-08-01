package main

import (
	"fmt"
)

// how modify struct attributes: functions must accept the
// pointer of the struct

func main() {

	my := MyStruct{}
	// it should be like this
	(&my).setA(5)
	// go allows you to use pointer methods without pointering
	my.setB(10)

	fmt.Println(my.A)
	fmt.Println(*my.B)

}

type MyStruct struct {
	A int
	B *int
}

func (my *MyStruct) setA(a int) {
	my.A = a
}

func (my *MyStruct) setB(b int) {
	// prevent panic: better than *my.B = b
	my.B = &b
}

