package main

import (
	"fmt"
)

var quadro string
var pquadro *string

func main() {

	cerchio := "cerchio"

	quadro = "quadro"
	pquadro = &quadro

	change(&cerchio)
	change(pquadro)

	fmt.Println(cerchio)
	fmt.Println(*pquadro)

	change2(&cerchio)
	change2(pquadro)

	fmt.Println(cerchio)
	fmt.Println(*pquadro)

	change3()
	fmt.Println(*pquadro)

	change4()
	fmt.Println(*pquadro)

}

// NEVER WORKS BECAUSE THE ARGUMENT IS ALWAYS PASSED AS A COPY (EVEN IF IT'S A POINTER!)
func change(s *string) {
	temp := "cubo"
	s = &temp
}

// ALWAYS WORKS BECAUSE THE POINTER VALUE IS NOT CHANGED. THE CONTENT OF THAT VALUE CHANGES UNDER THE HOOD
func change2(s *string) {
	if s == nil {
		return
	}
	// works because *s is the global value cerchio or quadro got from the copied pointer *s. Global values are always changeable
	*s = "piramide"
}

// ALWAYS WORK BECAUSE GLOBAL VARIABLE ARE CHANGEABLE FROM FUNCTIONS 
func change3() {
	quadro = "ipotenusa"
}

// ALWAYS WORK BECAUSE GLOBAL VARIABLE ARE CHANGEABLE FROM FUNCTIONS
func change4() {
	pquadro = &[]string{"cateto"}[0]
}

