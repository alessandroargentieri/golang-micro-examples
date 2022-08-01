// LA MUTABILITA' DEI PARAMETRI DI UNO STRUCT ATTRAVERSO I SUOI METODI (COMPORTAMENTI)

package main

import (
	"fmt"
)

type MyStruct struct {
	intero   int
	pstringa *string
}

func (m MyStruct) change1() { // non muta i parametri interni perche' m e' una copia
	m.intero = 5

	hi := "hi"
	m.pstringa = &hi
}

func (m MyStruct) change2() { // non muta i parametri interni perche' m e' una copia (non muta il valore nel puntatore pstringa - ma il suo contenuto a quella cella)
	m.intero = 5
	*m.pstringa = "hello"
}

func (m *MyStruct) change3() { // mutano sia il valore sia il puntatore perche' m e' passato come puntatore
	m.intero = 5

	hallo := "hallo"
	m.pstringa = &hallo
}

func (m *MyStruct) change4() { // mutano sia il valore sia il puntatore perche' m e' passato come puntatore
	m.intero = 5
	*m.pstringa = "bonjour"
}

func main() {
	ciao := "ciao"
	m := MyStruct{10, &ciao}

	fmt.Println("init value:", m.intero, "  ", *m.pstringa)

	m.change1()
	fmt.Println("change1   :", m.intero, "  ", *m.pstringa)

	m.change2()
	fmt.Println("change2   :", m.intero, "  ", *m.pstringa)

	m.change3() // qui change3() in realta' si riferisce ad un puntatore ma funziona lo stesso (lo crea Go internamente!)
	fmt.Println("change3   :", m.intero, "   ", *m.pstringa)

	m.change4() // qui change4() in realta' si riferisce ad un puntatore ma funziona lo stesso (lo crea Go internamente!)
	fmt.Println("change4   :", m.intero, "   ", *m.pstringa)

	// riproviamo gli ultimi due esplicitando che si tratta di puntatore a struct!
	fmt.Println()

	// reset initial value
	ciao = "ciao"
	m = MyStruct{10, &ciao}
	fmt.Println("init value:", m.intero, "  ", *m.pstringa)

	pm := &m

	pm.change3()
	fmt.Println("change3   :", m.intero, "   ", *m.pstringa)

	pm.change4()
	fmt.Println("change4   :", m.intero, "   ", *m.pstringa)

}

