package services

import (
	"errors"
	"fmt"
	"user-service/pkg/database"
	"user-service/pkg/model"
)

func AddUser(user model.User) (model.User, error) {
	if isUserExists(user.Username) {
		return model.User{}, errors.New("user already exists")
	}
	addUserSQL := "INSERT INTO users (Name, Email, Username, Password, Bio) VALUES (?, ?, ?, ?, ?)"
	_, err := database.DB.Exec(addUserSQL, user.Name, user.Email, user.Username, user.Password, user.Bio)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func isUserExists(username string) (userExists bool) {
	fmt.Println(username)
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE Username = ?)", username).Scan(&userExists)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return userExists
}
