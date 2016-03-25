package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName       string
	Email          string
	UnsafePassword string
	Password       []byte
}

func (db *DB) SubmitUser(u *User) {
	query := fmt.Sprintf(`INSERT INTO Users(username, password, email) VALUES('%s', '%s', '%s')`, u.UserName, u.Password, u.Email)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func (u *User) HashPassword() {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.UnsafePassword), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	u.Password = hashedPass
}
