package main

import (
	"log"
	"project_mini_golang/config"
	"project_mini_golang/controllers"
	"project_mini_golang/domain/repository"
	"project_mini_golang/domain/usecase"
	"project_mini_golang/middleware"
	"project_mini_golang/routes"
	"project_mini_golang/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	// Inisialisasi database
	DB, err := config.InitDB()
	if err != nil {
		log.Fatalf("Gagal menginisialisasi database: %v", err)
	}

	// database
	userRepo := repository.NewAuthRepository(DB)
	airQualityRepo := repository.NewAirQualityRepository(DB)
	alertRepo := repository.NewAlertRepository(DB)
	rekomendasiRepo := repository.NewRekomendasiRepository(DB)

	// service jwt
	secretKey := config.NewJWTConfig()
	jwtService := service.NewJWTService(secretKey)

	// usecase
	userUsecase := usecase.NewAuthUsecase(userRepo, jwtService)
	airQualityUsecase := usecase.NewAirQualityUsecase(airQualityRepo)
	alertUsecase := usecase.NewAlertUsecase(alertRepo)
	rekomendasiUsecase := usecase.NewRekomendasiUsecase(rekomendasiRepo)

	// controller
	authController := controllers.NewAuthController(userUsecase)
	airQualityController := controllers.NewAirQualityController(airQualityUsecase, alertUsecase, rekomendasiUsecase)
	alertController := controllers.NewAlertController(alertUsecase)
	rekomendasiController := controllers.NewRekomendasiController(rekomendasiUsecase)

	// middleware jwt
	jwtMiddleware := middleware.NewJWTMiddleware(secretKey)

	// routing autentikasi
	e := echo.New()
	authGroup := e.Group("")
	routes.AuthRoutes(authGroup, authController)

	// routing kualitas udara dengan middleware jwt
	route_group := e.Group("/airquality", jwtMiddleware.Handler)
	routes.AirqualityRoutes(route_group, airQualityController)
	routes.AlertRoutes(route_group, alertController)
	routes.RekomendasiRoutes(route_group, rekomendasiController)

	e.Start(":8000")
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal memuat file .env")
	}
}
