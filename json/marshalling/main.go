package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type student struct { 
  StudentId int `json:"id,required"` 
  LastName string `json:"lname"`  
  FirstName string `json:"fname"` 
  IsMarried bool `json:"-"` 
  IsEnrolled bool `json:"enrolled,omitempty"` 
  Courses []course `json:"classes"` 
} 

type course struct {
  Name string `json:"coursename,omitempty"`
  Number int `json:"coursenum,omitempty"`
  Hours int `json:"coursehours,omitempty"`
}

func main() {

  s := student{
    StudentId: 535, 
    LastName: "Ribezzo", 
    FirstName: "Annalucia", 
    IsMarried: false, 
    IsEnrolled: true, 
  } 

  c1 := course{Name: "World Lit", Number: 101, Hours: 3}
  //c2 := course{Name: "Intro to Go", Number: 101, Hours: 4}
  //c2 := course{Name: "", Number: 0, Hours: 0}
  c2 := course{}

  s.Courses = append(s.Courses, c1)
  s.Courses = append(s.Courses, c2)

  student, err := json.MarshalIndent(s, "", " ") 
  if err != nil { 
    fmt.Println(err) 
    os.Exit(1) 
  }
  fmt.Println(string(student)) 

}
