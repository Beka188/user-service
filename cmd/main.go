package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-service/internal/router"
	"user-service/pkg/database"
)

func main() {
	err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
	//users, err := services.ReadAllUsers()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(users)
	//fmt.Println(r)
	//err = services.DeleteUser("J@gmdail.com")
	//fmt.Println(err)
}
