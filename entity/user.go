package entity

import "time"

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password,omitempty"`
	Active   bool     `json:"active"`
	Meta     MetaData `json:"meta"`
}
type NewUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Active   bool   `json:"active"`
}

type Token struct {
	ID      string     `json:"id"`
	Value   string     `json:"value"`
	Active  bool       `json:"active"`
	User    Identifier `json:"user"`
	Expires time.Time  `json:"expires"`
	Meta    MetaData   `json:"meta"`
}

type NewToken struct {
	ID      string    `json:"id"`
	Value   string    `json:"value"`
	Active  bool      `json:"active"`
	UserID  string    `json:"user_id"`
	Expires time.Time `json:"expires"`
}

func CreateNewToken(userID string, expiration time.Time) NewToken {
	return NewToken{
		Active:  false,
		UserID:  userID,
		Expires: expiration,
	}
}

type MetaData struct {
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

type Identifier struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
