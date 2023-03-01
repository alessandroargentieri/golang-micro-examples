package main

import (
	"encoding/json"
	"strings"
	"errors"
	"fmt"
)

func main() {
	TestIsEqual()	
}


// IsEqual compares two values regardless of their types for testing purpose
func IsEqual(a, b interface{}) bool {
	A, _ := json.Marshal(a)
	B, _ := json.Marshal(b)

	return strings.ReplaceAll(string(A), "\"", "") == strings.ReplaceAll(string(B), "\"", "")
}

func TestIsEqual() {
	testName := "TestIsEqual"

	type testCase struct {
		description string
		inputA      interface{}
		inputB      interface{}
		expected    bool
	}

	testCases := []testCase{
		{
			description: "Comparing an integer and an int64",
			inputA:      3,
			inputB:      int64(3),
			expected:    true,
		},
		{
			description: "Comparing an unit16 and an int64",
			inputA:      uint16(3),
			inputB:      int64(3),
			expected:    true,
		},
		{
			description: "Comparing an integer in string format and an int64",
			inputA:      "3",
			inputB:      int64(3),
			expected:    true,
		},
		{
			description: "Comparing a pointer to a integer in string format and a pointer to an int64",
			inputA:      PointerTo("3"),
			inputB:      PointerTo(int64(3)),
			expected:    true,
		},
		{
			description: "Comparing a pointer to a float64 and a pointer to an int32",
			inputA:      PointerTo(float64(3)),
			inputB:      PointerTo(int32(3)),
			expected:    true,
		},
		{
			description: "Comparing a string and a pointer to a string",
			inputA:      "hello",
			inputB:      PointerTo("hello"),
			expected:    true,
		},
		{
			description: "Comparing two nil interfaces",
			inputA:      nil,
			inputB:      nil,
			expected:    true,
		},
		{
			description: "Comparing two errors",
			inputA:      fmt.Errorf("error!"),
			inputB:      errors.New("error!"),
			expected:    true,
		},
		{
			description: "Comparing a bool and a bool in string format",
			inputA:      true,
			inputB:      "true",
			expected:    true,
		},
		{
			description: "Comparing a pointer to a bool and a pointer to a bool in string format",
			inputA:      PointerTo(true),
			inputB:      PointerTo("true"),
			expected:    true,
		},
		{
			description: "Comparing a float in string format and a float",
			inputA:      "3.4",
			inputB:      float32(3.4),
			expected:    true,
		},
		{
			description: "Comparing a pointer to a float in string format and a pointer to a float",
			inputA:      PointerTo("3.4"),
			inputB:      PointerTo(float32(3.4)),
			expected:    true,
		},
		{
			description: "Comparing a slice of integers in string format and a slice of integers",
			inputA:      []string{"1", "2", "3"},
			inputB:      []int{1, 2, 3},
			expected:    true,
		},
		{
			description: "Comparing a slice of integers in string format and a slice of pointers to integers",
			inputA:      []string{"1", "2", "3"},
			inputB:      []*int{PointerTo(1), PointerTo(2), PointerTo(3)},
			expected:    true,
		},
		{
			description: "Comparing a map of pointers of integers in string format and a map of integers",
			inputA:      map[string]*string{"1": PointerTo("1"), "2": PointerTo("2")},
			inputB:      map[int]int{1: 1, 2: 2},
			expected:    true,
		},
		{
			description: "Comparing two different strings",
			inputA:      "hello",
			inputB:      "h1",
			expected:    false,
		},
		{
			description: "Comparing two different integers",
			inputA:      3,
			inputB:      2,
			expected:    false,
		},
		{
			description: "Comparing two different booleans",
			inputA:      true,
			inputB:      false,
			expected:    false,
		},
		{
			description: "Comparing two different slices",
			inputA:      []int{1, 2, 3},
			inputB:      []int{1, 2},
			expected:    false,
		},
	}
	for _, test := range testCases {
		if actual := IsEqual(test.inputA, test.inputB); actual != test.expected {
			fmt.Printf("%s failed (%s): expected %t, found %t\n", testName, test.description, test.expected, 
actual)
		} else {
		        fmt.Printf("%s succeeded (%s)\n", testName, test.description)
		}
	}
}


// PointerTo returns a pointer to a given value
func PointerTo[T any](value T) *T {
	return &value
}


