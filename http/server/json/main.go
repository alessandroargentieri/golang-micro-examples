package main

  import (
    "log"
    "net/http"
    "encoding/json"
  )

type Book struct {
  BookID   int    `json:"bookID,omitempty"`
  BookName string `json:"bookName,omitempty"`
  Author   string `json:"author,omitempty"`
}

func jsonPayload(w http.ResponseWriter, r *http.Request) {
  jsonDecoder := json.NewDecoder(r.Body)
  inputBook := Book{}
  err := jsonDecoder.Decode(&inputBook)
  if err != nil {
    log.Fatal(err)
  }
  jsonBytes, _ := json.Marshal(inputBook)
  log.Println(string(jsonBytes))
  w.Write([]byte(jsonBytes))
}

func main() {
  http.HandleFunc("/json", jsonPayload)            
  log.Fatal(http.ListenAndServe(":8080", nil))
}
