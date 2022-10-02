package webDemo

import (
	"fmt"
	"testing"
)

type User struct {
	Name     string
	Password string
	NickName string
}

func NewUser(name string, password string, nickName string) *User {
	return &User{Name: name, Password: password, NickName: nickName}
}

func (u *User) GetName() string {
	return u.Name
}

type Lcy struct {
	User
	Age int
}

func TestMainWebDemo2(t *testing.T) {
	// user := NewUser("lcy", "lcymm", "hh")
	// fmt.Printf("%T  value is %v\n", user, user)

	// fmt.Printf("name is %v\n", user.GetName())

	lcy := Lcy{User{"people", "friends", "stdout"}, 20}
	fmt.Printf("%T  value is %v\n", lcy, lcy)
}
