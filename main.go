package main

import (
	"bwastartup/app/campaign"
	"bwastartup/app/payment"
	"bwastartup/app/transaction"
	"bwastartup/app/user"
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/helper"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	webHandler "bwastartup/web/handler"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
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

	authService := auth.NewService()
	
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService)

	// Web CMS
	userWebHandler := webHandler.NewUserHandler(userService)
	campaignWebHandler := webHandler.NewCampaignHandler(campaignService, userService)

	router := gin.Default()
	router.Use(cors.Default())

	// load template pada direktori template
	router.HTMLRender = loadTemplates("./web/templates")

	// menambhakan routing utk gambar/file statis
	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	// grouping API by version
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checker", userHandler.IsEmailAvailable)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)
	
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transaction", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transaction/notification", transactionHandler.GetNotificationFromMidtrans)

	// Web CMS
	router.GET("/users", userWebHandler.Index)
	router.GET("/users/new", userWebHandler.New)
	router.POST("/users", userWebHandler.Create)
	router.GET("/users/edit/:id", userWebHandler.Edit)
	router.POST("/users/update/:id", userWebHandler.Update)
	router.GET("/users/avatar/:id", userWebHandler.NewAvatar)
	router.POST("/users/avatar/:id", userWebHandler.CreateAvatar)

	router.GET("/campaigns", campaignWebHandler.Index)
	router.GET("/campaigns/new", campaignWebHandler.New)
	router.POST("/campaigns", campaignWebHandler.Create)
	router.GET("/campaigns/image/:id", campaignWebHandler.NewImage)
	router.POST("/campaigns/image/:id", campaignWebHandler.CreateImage)

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


// template CMS
func loadTemplates(templatesDir string) multitemplate.Renderer {
  r := multitemplate.NewRenderer()

// load template pada direktori layouts
  layouts, err := filepath.Glob(templatesDir + "/layouts/*")
  if err != nil {
    panic(err.Error())
  }

// load semua template
  includes, err := filepath.Glob(templatesDir + "/**/*")
  if err != nil {
    panic(err.Error())
  }

  // Generate our templates map from our layouts/ and includes/ directories
  for _, include := range includes {
    layoutCopy := make([]string, len(layouts))
    copy(layoutCopy, layouts)
    files := append(layoutCopy, include)
    r.AddFromFiles(filepath.Base(include), files...)
  }
  return r
}