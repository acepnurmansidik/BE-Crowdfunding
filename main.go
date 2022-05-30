package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connection := "root:@tcp(127.0.0.1:3306)/crowfounding?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	//   jika ada error maka akan ditampilkan errornya
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// mencari user dengan email
	userByEmail, err := userRepository.FindByEmail("acepnurmansidik@gmail.com")
	if err != nil{
		fmt.Println(err.Error())
	}
	if (userByEmail.ID == 0){
		fmt.Println("accound not found")
	}else {
		fmt.Println(userByEmail)
	}
	// register user
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run()
	
}
