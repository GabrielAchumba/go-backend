package models

import "time"

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	IsDelete  bool      `json:"isDelete"`
}
