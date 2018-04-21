package thor

import (
	"fmt"
)

type User struct {
	User        string
	Password    string
	ID          int    `json:"user_id"`
	Token       string `json:"token"`
	PayPassword string
}

func NewUser(u, pw, ppw string) *User {
	return &User{
		User:        u,
		Password:    pw,
		PayPassword: ppw,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("User:%s, ID:%d, Token:%s", u.User, u.ID, u.Token)
}
