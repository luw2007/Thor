package res

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID          int    `json:"id"`
	User        string `json:"User"`
	Password    string `json:"password"`
	Token       string `json:"token"`
	PayPassword string `json:"pay_password"`
}

func NewUser(id int, name, password, payPassword string) *User {
	return &User{
		ID:          id,
		User:        name,
		Password:    password,
		PayPassword: payPassword,
	}
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *User) Type() Type {
	return ResUser
}

func (u *User) String() string {
	return fmt.Sprintf("User:%s, ID:%d, Token:%s", u.User, u.ID, u.Token)
}
