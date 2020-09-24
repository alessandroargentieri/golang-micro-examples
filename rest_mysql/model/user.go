package model

type User struct {
  Id      int    `json:"id,omitempty"`
  Name    string `json:"name,omitempty"`
  Email   string `json:"email,omitempty"`
}