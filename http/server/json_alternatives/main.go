package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// adding all the request headers in the response headers
		for key, values := range r.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		// setting hall other eaders and cookies
		// Keep in mind that headers should be set before writing the response body, otherwise the headers will not be sent to the client.
		w.Header().Add("X-Custom-Header", "value1")
		w.Header().Add("X-Custom-Header", "value2")
		cookie := http.Cookie{Name: "session_id", Value: "abc123", HttpOnly: true}
		http.SetCookie(w, &cookie)

		if r.Method == http.MethodPost ||
			r.Method == http.MethodPut ||
			r.Method == http.MethodPatch {

			var object interface{}
			err := json.NewDecoder(r.Body).Decode(&object)
			// ALTERNATIVE:
			/*
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "Error reading request body: "+err.Error(), http.StatusBadRequest)
					return
				}

				var object interface{}
				err = json.Unmarshal(body, &object)
			*/
			if err != nil {
				http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
				return
			}

			// alternative
			jsonBytes, _ := json.Marshal(object)
			fmt.Println(string(jsonBytes))
			// HTTP output suitable for []byte payloads: NEEDS a return AFTER else continues printing next w.Write statements
			w.Write([]byte(jsonBytes))
			return

			// ALTERNATIVE: suitable for string payloads, DOESN'T NEEDS a return AFTER
			// fmt.Fprint(w, string(jsonBytes))

		}

		fmt.Fprint(w, "hello world!")

	}))
	fmt.Println("HTTP server started on port 8080")
	http.ListenAndServe(":8080", nil)

}
