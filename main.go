package main

import (
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connection := "root:@tcp(127.0.0.1:3306)/crowfounding?charset=utf8mb4&parseTime=True&loc=Local"
  	// db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	// //   jika ada error maka akan ditampilkan errornya
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// // Jika koneksi berhasil
	// fmt.Println("Connecting success")

	// // buat variabel utk menampung data
	// var users []user.User

	// db.Find(&users)

	// for _, user:= range users{
	// 	fmt.Println("==========")
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// }

	router := gin.Default()
	router.GET("/users", handler)
	router.Run()
}

func handler(c *gin.Context){
	connection := "root:@tcp(127.0.0.1:3306)/crowfounding?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	//   jika ada error maka akan ditampilkan errornya
	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	// ubah ke dalam JSON
	c.JSON(200, users)
}