// COPY OBJECT WITHOUT KNOWING ITS TYPE
package main

import (
    "fmt"
    "reflect"
)

func main() {
   melone := &Melone{}
   melone2 := CopyGenericObject(*melone)
   fmt.Println(reflect.TypeOf(melone2))
}


func CopyGenericObject(input interface{}) interface{} {
    inputType := reflect.TypeOf(input)
    copy := reflect.New(inputType).Elem()
    copy.Set(reflect.ValueOf(input))
    return copy.Interface()
}

type Melone struct {}

