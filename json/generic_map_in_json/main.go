// JSON MARSHAL AND UNMARSHAL WITH A GENERIC MAP

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// ~~~~~~~~~ MARSHAL ~~~~~~~~~
	
	obj := Obj{
		Type: "type1",
		Name: "name1",
		Allocations: map[string][]string{
			"net1": []string{"A", "B", "C"},
			"net2": []string{"D", "E", "F"},
			"net3": []string{"G", "H", "I"},
		},
	}
	jsnBytes, _ := json.Marshal(obj)
	fmt.Println(string(jsnBytes))
	
	// ~~~~~~~~~ UNMARSHAL ~~~~~~~~~

	var obj2 Obj
	_ = json.Unmarshal(jsnBytes, &obj2)
	fmt.Println(obj2)

}

// STRUCT WITH A GENERIC MAP INSIDE
type Obj struct {
	Type        string              `json:"type"`
	Name        string              `json:"name"`
	Allocations map[string][]string `json:"allocations"`
}

