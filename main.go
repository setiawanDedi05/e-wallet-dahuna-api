package main

import (
	"fmt"
	"log"
	"mydiary_api/controller"
	"mydiary_api/database"
	"mydiary_api/middleware"
	"mydiary_api/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serverApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
	database.Database.AutoMigrate(&model.Wallet{})
	database.Database.AutoMigrate(&model.TransactionType{})
	database.Database.AutoMigrate(&model.DataPlan{})
	database.Database.AutoMigrate(&model.OperatorCard{})
	database.Database.AutoMigrate(&model.PaymentMethod{})
	database.Database.AutoMigrate(&model.Product{})
	database.Database.AutoMigrate(&model.Tip{})
	database.Database.AutoMigrate(&model.Transaction{})
	database.Database.AutoMigrate(&model.TransferHistory{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serverApplication() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	router.Static("/swaggerui/", "./swaggerui")

	publicRoutes := router.Group("/v1/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("v1/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	protectedRoutes.GET("/operator-card", controller.GetAllOperatorCards)
	protectedRoutes.POST("/operator-card", controller.AddOperatorCard)

	protectedRoutes.GET("/data-plan", controller.GetAllDataPlan)
	protectedRoutes.POST("/data-plan", controller.AddDataPlan)

	protectedRoutes.GET("/payment-method", controller.GetAllPaymentMethods)
	protectedRoutes.POST("/payment-method", controller.AddPaymentMethod)

	protectedRoutes.GET("/product", controller.GetAllProducts)
	protectedRoutes.POST("/product", controller.AddProduct)

	protectedRoutes.GET("/tips", controller.GetAllTip)
	protectedRoutes.POST("/tips", controller.AddTip)

	protectedRoutes.GET("/transaction-type", controller.GetAllTransactionType)
	protectedRoutes.POST("/transaction-type", controller.AddTransactionType)

	protectedRoutes.GET("/transaction", controller.GetAllTransacion)
	protectedRoutes.POST("/transaction", controller.AddTransaction)

	protectedRoutes.GET("/transfer-history", controller.GetAllTransferHistory)
	protectedRoutes.POST("/transfer-history", controller.AddTransferHistory)

	protectedRoutes.GET("/wallet", controller.GetAllWallet)
	protectedRoutes.POST("/wallet", controller.AddWallet)

	router.Run(":8888")
	fmt.Println("Server Running on port 8888")
}
