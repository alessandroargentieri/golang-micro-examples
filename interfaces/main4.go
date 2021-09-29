package main

import "fmt"

type Greeter interface {
   Greet()
}

type GreeterImpl struct {
   Wrapped string
}
func(g GreeterImpl) Greet() {
   fmt.Println(g.Wrapped)
}

var greeterGlobal = GreeterImpl{"Hello global wrapped!"}

var simpleVar string = "Hello simple var!"

func changeVar() {
   simpleVar = "simpleVar changed again!"
}

func changeGivenVar(variable string) {
   variable = "This given variable has changed!"
}
func main() {
   greeter := GreeterImpl{ "Hello wrapped!" }
   greeter.Greet()
   greeter.Wrapped = "Hello changed wrapped"
   greeter.Greet()

   fmt.Println("----------")
   greeterGlobal.Greet()
   greeterGlobal.Wrapped = "Global wrapped has changed!"
   greeterGlobal.Greet()

   fmt.Println("---------")
   fmt.Println(simpleVar)
   simpleVar = "simpleVar has changed!"
   fmt.Println(simpleVar)
   changeVar()
   fmt.Println(simpleVar)

   fmt.Println("--------")
   variable := "This is a var"
   fmt.Println(variable)
   changeGivenVar(variable)
   fmt.Println(variable)
   
}

