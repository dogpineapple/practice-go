package models

// represents the `users` table
// uppercase means it is public or exported and can be accessed by other packages.
type User struct {
	ID int64 `json:"id"`
  Username string `json:"username"`
  Ign string `json:"ign"` 
  World string `json:"world"`
}