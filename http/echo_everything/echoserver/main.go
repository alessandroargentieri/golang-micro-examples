package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Print the verb
		response := fmt.Sprintf("HTTP Verb: %s\n", r.Method)

		// Print the headers
		reqHeaders := "Headers:\n"
		for key, values := range r.Header {
			for _, value := range values {
				reqHeaders = reqHeaders + fmt.Sprintf("  %s: %s\n", key, value)
			}
		}
		response = response + reqHeaders

		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// Print the request body
		response = response + fmt.Sprintf("Body: %s\n", string(body))

		// Respond with the received information
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	// Start the server
	fmt.Println("Server is listening on :8585...")
	http.ListenAndServe(":8585", nil)
}
