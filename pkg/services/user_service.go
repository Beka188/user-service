package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"user-service/pkg/database"
	"user-service/pkg/model"
)

func CreateUser(user model.User) (model.User, error) {
	hashedPassword, err := HashPassword(user.Password)
	addUserSQL := "INSERT INTO users (Name, Email, Username, Password, ProfilePicture, Bio) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := database.DB.Exec(addUserSQL, user.Name, user.Email, user.Username, hashedPassword, user.ProfilePicture, user.Bio)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
			return model.User{}, ErrConflict
		}
		return model.User{}, err
	}
	user.Password = hashedPassword
	user.ID, _ = result.LastInsertId()
	return user, nil
}

func ReadPublicUser(id int) (model.UserPublic, error) {
	user := model.UserPublic{}
	getUserSQL := "SELECT ID, Name, Username, ProfilePicture, Bio, CreatedAt, UpdatedAt FROM users WHERE ID = ?"
	err := database.DB.QueryRow(getUserSQL, id).Scan(&user.ID, &user.Name, &user.Username, &user.ProfilePicture, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserPublic{}, ErrNotFound
		}
		return model.UserPublic{}, err
	}
	return user, nil
}

func ReadUser(id int) (model.User, error) {
	user := model.User{}
	getUserSQL := "SELECT ID, Name, Email, Username, Password, ProfilePicture, Bio, CreatedAt, UpdatedAt FROM users WHERE ID = ?"
	err := database.DB.QueryRow(getUserSQL, id).Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Password, &user.ProfilePicture, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, ErrNotFound
		}
		return model.User{}, err
	}
	fmt.Println("OLD")
	fmt.Println(user)
	return user, nil
}

func ReadAllUsers() ([]model.UserPublic, error) {
	var users []model.UserPublic
	getUserSQL := "SELECT ID, Name, Username, ProfilePicture, Bio, CreatedAt, UpdatedAt FROM users"
	rows, err := database.DB.Query(getUserSQL)
	if err != nil {
		return users, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var user model.UserPublic
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.ProfilePicture, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func UpdateUser(user model.User, id int) (model.User, error) {
	updateUserSQL := "UPDATE users SET Name = ?, Username = ?, Password = ?, ProfilePicture = ?, Bio = ? WHERE ID = ?"
	_, err := database.DB.Exec(updateUserSQL, user.Name, user.Username, user.Password, user.ProfilePicture, user.Bio, id)
	if err != nil {
		fmt.Println("Here ")
		fmt.Println(err)
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
			return model.User{}, ErrConflict
		}
		return model.User{}, err
	}
	fmt.Println("HEHEHHEHE")
	_, err = ReadPublicUser(id)
	fmt.Println(err)
	if errors.Is(err, ErrNotFound) {
		return model.User{}, ErrNotFound
	}
	return user, nil
}

func DeleteUser(id int) error {
	deleteUserSQL := "DELETE FROM users WHERE ID = ?"
	result, err := database.DB.Exec(deleteUserSQL, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrNotFound
	} else {
		return nil
	}
}
