package main

import (
   "fmt"
   "database/sql"
    _ "github.com/go-sql-driver/mysql"
)   

// INTERFACE **********************************************************
type Database interface {
   Open(configs map[string]string) (Database, error)
   Query(query string) (interface{}, error)
   Close()
}

// IMPLEMENTED DB ****************************************************
type DatabaseImpl struct {
   Db *sql.DB  
}
func(d DatabaseImpl) Open(configs map[string]string) (Database, error) {
   db, err := sql.Open(configs["dbms"], fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs["username"], configs["password"], configs["host"], configs["port"], configs["dbname"]))
   defer db.Close()
   d.Db = db
   return d, err
}
func(d DatabaseImpl) Close() {
   d.Db.Close()
}
func(d DatabaseImpl) Query(query string) (interface{}, error) {
   result, err := d.Db.Query(query)
   defer result.Close()
   return result, err
}

// IMPLEMENTED MOCK **************************************************
type DatabaseMock struct {
  outputMap map[string]interface{} 
}
func(d DatabaseMock) Open(configs map[string]string) (Database, error) {
   return d, d.outputMap["openError"].(error)
}
func(d DatabaseMock) Close() {}
func(d DatabaseMock) Query(query string) (interface{}, error) {
   mockedOutput := d.outputMap["query"]
   switch mockedOutput.(type) {
      case error:
          return nil, mockedOutput.(error)   
      default:
          return mockedOutput, nil
   }
}
func(d DatabaseMock) SetQueryOutput(query string, output interface{}) {
   d.outputMap[query] = output
}
func(d DatabaseMock) SetOpenConnectionError(openError error) {
   d.outputMap["openError"] = openError
}


func main() {
  fmt.Println("Hello main!")
}
