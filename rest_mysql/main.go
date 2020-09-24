package main

import (
	database "goql/dao"
	handlers "goql/handlers"
  "log"
  "net/http"
)


func main() {
	database.Connect()

	http.HandleFunc("/add", handlers.AddUser)             // curl -s -D - -H "Content-Type: application/json;charset=utf-8" -X POST http://localhost:8080/json -d '{"bookID": 3232, "bookName": "Cime Tempestose", "author": "Gigi Bom"}'
  http.HandleFunc("/get", handlers.GetAllUsers)           // curl -H "Authorization: superSecretToken" http://localhost:8080/token

  log.Fatal(http.ListenAndServe(":8080", nil))

}