#! /bin/bash

go run $(cat > main.go << 'EOF'
package main

import (
  "net/http"
  "log"
)

func main() {
   http.HandleFunc("/hello",  func(w http.ResponseWriter, r *http.Request){
     w.Write([]byte("Hello world!"))
   })
   log.Fatal(http.ListenAndServe(":8080", nil))
   log.Println("HTTP Server inited on port 8080")
}
EOF
ls -t | head -1) &

