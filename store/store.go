package store

import (
	"database/sql"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func AddUser(email, password string) {
	user := User{
		Email:    email,
		Password: password,
	}
	users = append(users, user)
}

func FindUser(email, password string) (user User, err error) {
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}

	return User{}, sql.ErrNoRows
}

func Users() (users []User) {
	return users
}
