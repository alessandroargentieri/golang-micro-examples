package main

  import (
    "log"
    "net/http"
  )

// get the method
func getMethod(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost { 
      w.Write([]byte("It's a POST request"))
      return
  } else if r.Method == http.MethodGet {
      w.Write([]byte("It's a GET request"))
      return
  }
  w.Write([]byte("It's not a POST nor a GET request"))  
}

// check the headers
func checkHeaders(w http.ResponseWriter, r *http.Request) {
  auth := r.Header.Get("Authorization")
  if auth != "superSecretToken" {
    w.WriteHeader(http.StatusUnauthorized)
    w.Write([]byte("Authorization token not recognized"))
    return
  }
  msg := "hello client!"
  w.Write([]byte(msg))
}

func main() {
  http.HandleFunc("/getmethod", getMethod)
  http.HandleFunc("/token", checkHeaders)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
