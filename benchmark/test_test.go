package main

import (
	"testing"

	"github.com/alessandroargentieri/gofunk"
)

type MyStruct struct {
	A int
	B string
	C bool
	D []int
	E []string
}

func f1(input interface{}) (interface{}, error) {
	in := input.(MyStruct)
	in.A++
	in.B = in.B + "-"
	in.C = !in.C
	in.D = append(in.D, in.A)
	in.E = append(in.E, in.B)
	return in, nil
}

func f2(in MyStruct) (MyStruct, error) {
	in.A++
	in.B = in.B + "-"
	in.C = !in.C
	in.D = append(in.D, in.A)
	in.E = append(in.E, in.B)
	return in, nil
}

func BenchmarkWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Measure time taken for some operation here
		_ = gofunk.OptionalOf(MyStruct{A: 0, B: "hello", C: false, D: []int{0}, E: []string{"hello"}}).
			Map(f1).
			Map(f1).
			Map(f1).
			Map(f1).
			Map(f1).
			Map(f1).
			Map(f1).
			Map(f1).
			Map(f1).
			Map(f1).
			Get().(MyStruct)
	}

}

func BenchmarkWithout(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Measure time taken for some operation here
		myStruct := MyStruct{A: 0, B: "hello", C: false, D: []int{0}, E: []string{"hello"}}
		var err error
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		myStruct, err = f2(myStruct)
		if err != nil {
			// nothing
		}
		_ = myStruct
	}
}
