package main

import (
	"fmt"
	"log"
	"user-service/pkg/database"
	"user-service/pkg/services"
)

func main() {
	err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	//newUser := model.User{Name: "Beka", Email: "J@gmail.com", Username: "Kira23", Password: "123", ProfilePicture: ";;", Bio: "Smth about me"}
	//_, err = services.AddUser(newUser)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//newNewUser := model.User{Name: "Heheheheh", Email: "J@com", Username: "Kira23", Password: "12345", ProfilePicture: "KIRA", Bio: "Justice"}
	//
	//_, err = services.UpdateUser(newNewUser, "J@gmail.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	_, err = services.GetPublicUser("Kira234")
	fmt.Println(err)
}
