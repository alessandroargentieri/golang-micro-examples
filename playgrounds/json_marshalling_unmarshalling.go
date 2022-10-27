
// YOU CAN RETURN A POINTER THROUGH AN INTERFACE, CAST IT TO A POINTER AGAIN AND FETCH THE VALUE WITH NO PROBLEMS
package main

import (
	"fmt"
	"encoding/json"
)


type Car struct{
    Model *string `json:"model,omitempty"`
    Brand *string `json:"brand,omitempty"`
}

func pointerToString(s string) *string{
   return &s
}

func main() {

        // pointer of a struct with pointers inside
	car := Car{
	   Model: pointerToString("Delta"),
	   Brand: pointerToString("Lancia"),
	}
	pcar := &car
	
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// THIS SHOWS HOW MARSHALLING IN JSON A POINTER OR A STRUCT BRINGS TO THE SAME RESULT!
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~	
	bytes1, _ := json.Marshal(pcar)
	bytes2, _ := json.Marshal(car)
	fmt.Println(string(bytes1))
	fmt.Println(string(bytes2))
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	
	fmt.Println()
	
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// UNMARSHALLING IN WHICH WE KNOW THE TYPE AND WE PASS THE POINTER OF THE EMPTY STRUCT
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	jsnstr := string(bytes1)
	var c Car
	_ = json.Unmarshal([]byte(jsnstr), &c)
	
	fmt.Printf("Model: %s, Brand: %s \n", *c.Model, *c.Brand)
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	
	fmt.Println()
	
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// UNMARSHALLING IN WHICH WE KNOW THE TYPE AND WE PASS THE POINTER OF AN EMPTY INTERFACE:
	// HERE WE NOTICE HOW THE EMPTY INTERFACE WILL CONTAINS IN REALITY A map[string]interface{}
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	var generic1 interface{}
	_ = json.Unmarshal(bytes1, &generic1)
	fmt.Printf("%+v \n", generic1)
	fmt.Printf("%T \n", generic1)
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	
	fmt.Println()
	
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// UNMARSHALLING IN WHICH WE KNOW THE TYPE AND WE PASS THE POINTER OF THE EMPTY map[string]interface{}
	// IT'S EXACTLY THE SAME AS THE EXAMPLE ABOVE IN WHICH WE PASSED AN interface{}
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	var generic2 map[string]interface{}
	_ = json.Unmarshal(bytes1, &generic2)
	fmt.Printf("%+v \n", generic2)
	fmt.Printf("%T \n", generic2)
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	
	fmt.Println()
	
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// IN FACT, IF WE PRINT THE RESULT OF CASTING THE interface{} TO THE map[string]interface{} WE GET TRUE!
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	_, ok := generic1.(map[string]interface{})
	fmt.Println(ok) 
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	
	fmt.Println()
	
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// ONCE WE UNMARSHAL A JSON INTO A map[string]interface{} IS NOT POSSIBLE TO CAST TO THE SPECIFIC TYPE!
 	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	_, ok = generic1.(Car)
	fmt.Println(ok) 
	
	//_, ok = generic2.(Car) returns an error because only interface{} can be cast to other types, not map[string]interface{} !!!!!!!!
	_, ok = generalizeMe(generic2).(Car)
	fmt.Println(ok)
	
	// FINAL CONSIDERATION: fmt.Printf("%T \n", value) TELLS ALWAYS THE TRUTH ABOUT THE REAL TIPE OF THE VALUE.. EVEN IF YOU HAVE AN INTERFACE!!!!
	var s = "ciao"
	var si interface{} = s
	fmt.Printf("%T %T \n", s, si)
        
}

func generalizeMe(i interface{}) interface{} {
   return i
}
