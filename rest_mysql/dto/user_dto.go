package dto

type UserDto struct {
  Name  string `json:"name,omitempty"`
  Email string `json:"email,omitempty"`
}