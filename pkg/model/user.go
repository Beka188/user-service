package model

import "time"

type User struct {
	ID             int64
	Name           string
	Email          string
	Username       string
	Password       string
	ProfilePicture string
	Bio            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UserPublic struct {
	ID             int
	Name           string
	Username       string
	ProfilePicture string
	Bio            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
