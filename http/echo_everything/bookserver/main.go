package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		type Book struct {
			Title  string `json:"title,omitempty"`
			Author string `json:"author,omitempty"`
		}

		fmt.Println("Headers:")
		for k, v := range r.Header {
			fmt.Printf("  %s: %s\n", k, v)
			if k == "Authorization" {
				if strings.HasPrefix(v[0], "Bearer") {
					fmt.Printf("Bearer authentication used with token %s\n", strings.Split(v[0], "Bearer ")[1])
				} else {
					encodedCred := strings.Split(v[0], "Basic ")[1]
					decodedCreds, _ := base64.StdEncoding.DecodeString(encodedCred)
					credentials := strings.Split(string(decodedCreds), ":")
					fmt.Printf("(Basic authentication used with user %s and password %s)\n", credentials[0], credentials[1])
				}
			}
		}

		book := Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		jsonRespBytes, _ := json.Marshal(book)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRespBytes)
	})

	// Start the server
	fmt.Println("Book server is listening on :8585...")
	http.ListenAndServe(":8585", nil)
}
