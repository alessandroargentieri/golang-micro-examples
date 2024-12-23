package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	name, ok := queryString["name"]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing name"))
		return
	}
	msg := "Hello " + name[0]
	w.Write([]byte(msg))
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hi", hiHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)

}
