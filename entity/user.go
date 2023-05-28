package entity

import "fmt"

type User struct {
	ID       int
	Username string
	Password string
}

func (u *Product) UserDetail() {
	fmt.Println("ID: ", u.ID)
	fmt.Println("username: ", u.Name)
	fmt.Println("")
}
