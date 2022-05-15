package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connection := "root:@tcp(127.0.0.1:3306)/crowfounding?charset=utf8mb4&parseTime=True&loc=Local"
  	_, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connecting success")
}