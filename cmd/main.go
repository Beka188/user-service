package main

import (
	"fmt"
	"log"
	"user-service/pkg/database"
	"user-service/pkg/model"
	"user-service/pkg/services"
)

func main() {
	err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	newUser := model.User{Name: "Beka", Email: "J@gmail.com", Username: "Lawliet", Password: "123", ProfilePicture: ";;", Bio: "Smth about me"}
	user, err := services.AddUser(newUser)
	if err != nil {
		return
	}
	fmt.Println(user)
}
