package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
type NewUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Active   bool   `json:"active"`
}

type Token struct {
	ID        string    `json:"id"`
	Value     string    `json:"value"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	Username  string    `json:"username"`
	UserID    string    `json:"-"`
	Expires   time.Time `json:"expires"`
}
