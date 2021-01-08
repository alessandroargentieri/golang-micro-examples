package main

  import (
    "log"
    "reflect"
    "net/http"
    "encoding/json"
    "bytes"
  )

type Book struct {
  BookID   int    `json:"bookID,omitempty"`
  BookName string `json:"bookName,omitempty"`
  Author   string `json:"author,omitempty"`
}

func jsonPayload(w http.ResponseWriter, r *http.Request) {
//  jsonDecoder := json.NewDecoder(r.Body)
  inputBook := Book{}
//  err := jsonDecoder.Decode(&inputBook)
    log.Println(reflect.TypeOf(r.Body))
    buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
    err := json.Unmarshal([]byte(buf.String()), &inputBook)
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
