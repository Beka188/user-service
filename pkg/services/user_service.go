package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"user-service/pkg/database"
	"user-service/pkg/model"
)

func AddUser(user model.User) (model.User, error) {
	hashedPassword, err := HashPassword(user.Password)
	addUserSQL := "INSERT INTO users (Name, Email, Username, Password, ProfilePicture, Bio) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = database.DB.Exec(addUserSQL, user.Name, user.Email, user.Username, hashedPassword, user.ProfilePicture, user.Bio)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
			return model.User{}, fmt.Errorf("user with username or/and email already exists: %s %s ", user.Username, user.Email)
		}
		return model.User{}, err
	}
	return user, nil
}

func GetPublicUser(username string) (model.PublicUser, error) {
	user := model.PublicUser{}
	getUserSQL := "SELECT Name, Username, ProfilePicture, Bio, CreatedAt, UpdatedAt FROM users WHERE Username = ?"
	err := database.DB.QueryRow(getUserSQL, username).Scan(&user.Name, &user.Username, &user.ProfilePicture, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.PublicUser{}, errors.New("user not found: " + username)
		}
		return model.PublicUser{}, err
	}
	return user, nil
}

func GetUser(username string) (model.User, error) {
	user := model.User{}
	getUserSQL := "SELECT Name, Email, Username, Password, ProfilePicture, Bio, CreatedAt, UpdatedAt FROM users WHERE Username = ?"
	err := database.DB.QueryRow(getUserSQL, username).Scan(&user.Name, &user.Email, &user.Username, &user.Password, &user.ProfilePicture, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, errors.New("user not found: " + username)
		}
		return model.User{}, err
	}
	return user, nil
}

func UpdateUser(user model.User, oldEmail string) (model.User, error) {
	updateUserSQL := "UPDATE users SET Name = ?, Username = ?, Password = ?, ProfilePicture = ?, Bio = ? WHERE Email = ?"
	_, err := database.DB.Exec(updateUserSQL, user.Name, user.Username, user.Password, user.ProfilePicture, user.Bio, oldEmail)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
			return model.User{}, fmt.Errorf("user with username already exists: %s", user.Username)
		}
		return model.User{}, err
	}
	return user, nil
}

//
//func DeleteUser(username string) error {
//
//}

func isUserExists(username string) (userExists bool) {
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE Username = ?)", username).Scan(&userExists)
	if err != nil {
		return false
	}
	return userExists
}
