package store

import (
	"database/sql"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

var users []User

func UpdateUserSecret(email, secret string) {
	for i, user := range users {
		if user.Email == email {
			users[i].Secret = secret
		}
	}
}

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

func FindUserByEmail(email string) (user User, err error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}

	return User{}, sql.ErrNoRows
}

func Users() []User {
	return users
}
