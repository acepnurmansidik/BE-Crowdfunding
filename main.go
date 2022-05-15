package main

import (
	"bwastartup/user"
	"fmt"
	"log"

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
	// Jika koneksi berhasil
	fmt.Println("Connecting success")

	// buat variabel utk menampung data
	var users []user.User

	db.Find(&users)

	for _, user:= range users{
		fmt.Println("==========")
		fmt.Println(user.Name)
		fmt.Println(user.Email)
	}

}