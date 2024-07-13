package model

import "time"

type User struct {
	Name           string
	Email          string
	Username       string
	Password       string
	ProfilePicture string
	Bio            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type PublicUser struct {
	Name           string
	Username       string
	ProfilePicture string
	Bio            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
