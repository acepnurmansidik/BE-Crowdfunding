package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaigns, err := campaignService.FindCampaign(7)
	if err != nil {
		fmt.Println("Not found")
	}

	fmt.Println(campaigns)
	fmt.Println("#########################")

	campaignAll, err := campaignRepository.FindAll()
	
	fmt.Println(campaignAll)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checker", userHandler.IsEmailAvailable)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
	
}

// buat middleware
// bungkus menjadi function utk mengembalikan sebuah func gin handler
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	// return func gin handler
	return func (c *gin.Context){
		// ambil token dari header
		authHeader := c.GetHeader("Authorization")

		// cek isinya ada bearer
		if !strings.Contains(authHeader, "Bearer"){
			// buat responsenya
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// AbortWithStatusJSON, berfungsi utk menghentikan program jika ada error
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		stringToken := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			// set token ke var string
			stringToken = arrayToken[1]
		}

		// validasi token
		token, err := authService.ValidateToken(stringToken)
		if err != nil {
			// buat responsenya
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// AbortWithStatusJSON, berfungsi utk menghentikan program jika ada error
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ubah token ke jwt mapClaims
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			// buat responsenya
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// AbortWithStatusJSON, berfungsi utk menghentikan program jika ada error
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ambil user_id di dapat dari client, lalu ubah ke dalam int
		userID := int(claim["user_id"].(float64))

		// cari user yg ddapat dari token yang dikirim dari client
		user, err := userService.GetUserByID(userID)
		if err != nil {
			// buat responsenya
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// AbortWithStatusJSON, berfungsi utk menghentikan program jika ada error
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set user yang sedag akses aplikasi
		c.Set("currentUser", user)
	}
}
