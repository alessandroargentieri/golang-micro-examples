// SLICE TYPE MODIFIABLE VIA MUTATOR METHOD

package main

import (
	"fmt"
)

type APIConditions []string

func (c *APIConditions) Apply(items ...string) {
   if c == nil {
      return
   }
   *c = append(*c, items...) 
}

func main() {
   var conditions APIConditions = make([]string, 0)
   fmt.Println(conditions)

   conditions.Apply("ciao", "hello")
   fmt.Println(conditions)
}

