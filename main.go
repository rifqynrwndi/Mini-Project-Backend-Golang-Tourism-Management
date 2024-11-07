package main

import (
	"log"
	"tourism-monitoring/config"
	authController "tourism-monitoring/controllers/auth"
	touristsController "tourism-monitoring/controllers/tourists"
	placesController "tourism-monitoring/controllers/places"
	"tourism-monitoring/middleware"
	authRepo "tourism-monitoring/repositories/auth"
	touristsRepo "tourism-monitoring/repositories/tourists"
	placesRepo "tourism-monitoring/repositories/places"
	"tourism-monitoring/routes"
	authService "tourism-monitoring/services"
	touristsService "tourism-monitoring/services/tourists"
	placesService "tourism-monitoring/services/places"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	config.MigrateDB(db)

	e := echo.New()

	// Initialize Auth
	authJwt := middleware.JwtTourism{}
	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	// Initialize Tourists
	touristsRepo := touristsRepo.NewTouristsRepo(db)
	touristsService := touristsService.NewTouristsService(touristsRepo)
	touristsController := touristsController.NewTouristsController(touristsService)

	// Initialize Places
	placesRepo := placesRepo.NewPlacesRepo(db)
	placesService := placesService.NewPlacesService(placesRepo)
	placesController := placesController.NewPlacesController(placesService)

	routeController := routes.RouteController{
		AuthController:     authController,
		TouristsController: touristsController,
		PlacesController:   placesController,
	}
	routeController.InitRoute(e)

	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
