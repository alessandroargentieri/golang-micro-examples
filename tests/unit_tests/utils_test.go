package main

import (
  "time"
  "testing"
)

// testing public func
func TestSumIntsCorrect(t *testing.T) {
  expected := 10
  actual := SumInts(5, 5)
  if actual != expected {
    t.Errorf("Sum was incorrect")
  }
}

func TestSumIntsIncorrect(t *testing.T) {
  expected := 10
  actual := SumInts(5, 4)
  if actual != expected { 
    t.Errorf("Sum was incorrect")
  }
}

// testing private func
func TestSubIntsCorrect(t *testing.T) {
  expected := 0
  actual := subInts(5, 5)
  if actual != expected { 
    t.Errorf("Sub was incorrect")
  }
}

func TestSubIntsIncorrect(t *testing.T) {
  expected := 0
  actual := subInts(5, 4)
  if actual != expected { 
    t.Errorf("Sub was incorrect")
  }
}


// test tables
func TestSumIntsTables(t *testing.T) {
	tables := []struct {
		input1 int
		input2 int
		result int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 8},
	}

	for _, table := range tables {
		total := SumInts(table.input1, table.input2)
		if total != table.result {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.input1, table.input2, total, table.result)
		}
	}
}

// get variables from tested class
func TestGetVariable(t *testing.T) {
   if variable!="VARIABLE" {
      t.Errorf("Local variable has not the expected value")
   }
}

// access, modify and verify pointers from tested class
func TestGetAndChangePointerValue(t *testing.T) {
   if pointer == nil {
      t.Errorf("Cannot access pointer")
   } else if *pointer != "VALUE" {
      t.Errorf("Pointer has not the expected value")
   }
   *pointer = "value"
   actual := getPointerValue()
   if actual != "value" {
     t.Errorf("Cannot modify pointer value from outside")
   }
}



// testing struct function
func TestGetModelFromRamStruct(t *testing.T) {
  ram := Ram{
    Model: "DDR SDRAM",
    Year:  time.Date(2019, 6, 5, 11, 35, 04, 0, time.UTC),  //for a new line we need a comma here!
  }

  expected := "DDR SDRAM 2019-06-05 11:35:04 +0000 UTC"
  actual := ram.GetModel()
  if actual != expected { 
    t.Errorf("The method GetModel() didn't return what expected!")
  }
}


