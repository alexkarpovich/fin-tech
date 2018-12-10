package entity

import "encoding/json"

type (
	// User struct represents general User model
	User struct {
		ID        int64    `json:"id"`
		Email     string   `json:"email"`
		FirstName string   `json:"first_name"`
		LastName  string   `json:"last_name"`
		CreatedAt Time     `json:"created_at"`
		DeletedAt NullTime `json:"deleted_at"`
	}
)

func (u User) Serialize(group string) ([]byte, error) {
	return json.Marshal(&struct {
		ID int64
	}{
		ID: u.ID,
	})
}
