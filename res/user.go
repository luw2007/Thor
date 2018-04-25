package res

import (
	"encoding/json"
	"fmt"
)

type user struct {
	ID          int    `json:"id"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Token       string `json:"token"`
	PayPassword string `json:"pay_password"`
}

func NewUser(id int, name, password, payPassword string) *user {
	return &user{
		ID:          id,
		User:        name,
		Password:    password,
		PayPassword: payPassword,
	}
}

func (u *user) GetID() int {
	return u.ID
}

func (u *user) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *user) Type() Type {
	return User
}

func (u *user) String() string {
	return fmt.Sprintf("user:%s, ID:%d, Token:%s", u.User, u.ID, u.Token)
}
