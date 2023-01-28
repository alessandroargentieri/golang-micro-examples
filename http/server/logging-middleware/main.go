package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request headers
		log.Println("Request Headers:", r.Header)

		// Log request body
		body := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(body)
		log.Println("Request Body:", string(body))

		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()

		// Log response headers
		log.Println("Response Headers:", w.Header())

		// Log response body
		body = []byte{}
		if res, ok := w.(responseBody); ok {
			body = res.body
		}
		log.Println("Response Body:", string(body))

		// Log response time
		log.Println("Response Time:", end.Sub(start))
	})
}

type responseBody struct {
	http.ResponseWriter
	body []byte
}

func (w responseBody) Write(b []byte) (int, error) {
	w.body = b
	return w.ResponseWriter.Write(b)
}

func main() {
	http.Handle("/", loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

			//
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body: "+err.Error(), http.StatusBadRequest)
				return
			}

			var object interface{}
			err = json.Unmarshal(body, &object)
			//
			//var body interface{}
			//err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				//http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
				//return

				w.Write([]byte(err.Error()))
				return
			}

			w.Write([]byte("Ciao"))
			return
			// rewrite received json body in output

			// alternative
			jsonBytes, _ := json.Marshal(object)
			log.Println(string(jsonBytes))
			w.Write([]byte(jsonBytes))
			return

			// alternative
			//	   fmt.Fprint(w, string(jsonBytes))
			//	   return

		}

		fmt.Fprint(w, "hello world!")

	})))
	http.ListenAndServe(":8080", nil)
	fmt.Println("HTTP server started on port 8080")
}
