#! /bin/bash

cat > main.go << 'EOF'
package main

import (
  "net/http"
  "log"
  "fmt"
)

func main() {
   http.HandleFunc("/hello",  func(w http.ResponseWriter, r *http.Request){
     queryString := r.URL.Query()
     name := queryString.Get("name")
     if name == "" {
         name = "world"
     }
     w.Write([]byte(fmt.Sprintf("Hello %s!", name)))
   })
   log.Fatal(http.ListenAndServe(":8080", nil))
   log.Println("HTTP Server inited on port 8080")
}
EOF

docker run -it -d --rm -v $PWD:/usr/src/myapp -p 8080:8080 -w /usr/src/myapp golang:1.17 go run main.go
docker run -it -d -p 8080:8080 --name localtunnel efrecon/localtunnel --port 8080
