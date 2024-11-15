package main

import (
	"log"
	"tourism-monitoring/config"
	AIController "tourism-monitoring/controllers/ai"
	authController "tourism-monitoring/controllers/auth"
	placesController "tourism-monitoring/controllers/places"
	touristsController "tourism-monitoring/controllers/tourists"
	trashReportController "tourism-monitoring/controllers/trash_report"
	visitReportController "tourism-monitoring/controllers/visit_report"
	weatherController "tourism-monitoring/controllers/weather"
	

	"tourism-monitoring/middleware"
	authRepo "tourism-monitoring/repositories/auth"
	placesRepo "tourism-monitoring/repositories/places"
	touristsRepo "tourism-monitoring/repositories/tourists"
	trashReportRepo "tourism-monitoring/repositories/trash_report"
	visitReportRepo "tourism-monitoring/repositories/visit_report"
	weatherRepo "tourism-monitoring/repositories/weather"

	"tourism-monitoring/routes"
	authService "tourism-monitoring/services"
	AIService "tourism-monitoring/services/ai"
	placesService "tourism-monitoring/services/places"
	touristsService "tourism-monitoring/services/tourists"
	trashReportService "tourism-monitoring/services/trash_report"
	visitReportService "tourism-monitoring/services/visit_report"
	weatherService "tourism-monitoring/services/weather"

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

	// Initialize Visit Report
	visitReportRepo := visitReportRepo.NewVisitReportRepo(db)
	visitReportService := visitReportService.NewVisitReportService(visitReportRepo)
	visitReportController := visitReportController.NewVisitReportController(visitReportService)

	// Initialize Trash Report
	trashReportRepo := trashReportRepo.NewTrashReportRepo(db)
	trashReportService := trashReportService.NewTrashReportService(trashReportRepo)
	trashReportController := trashReportController.NewTrashReportController(trashReportService)

	// Initialize AI
	aiService, err := AIService.NewAIService(visitReportRepo, placesRepo)
	if err != nil {
		log.Fatalf("AI service initialization failed: %v", err)
	}
	aiController := AIController.NewAIController(aiService)

	// Initialize Weather
	weatherRepo := weatherRepo.NewWeatherRepo()
	weatherService := weatherService.NewWeatherService(weatherRepo)
	weatherController := weatherController.NewWeatherController(weatherService)
	

	routeController := routes.RouteController{
		AuthController:        authController,
		TouristsController:    touristsController,
		PlacesController:      placesController,
		VisitReportController: visitReportController,
		TrashReportController: trashReportController,
		AIController:          aiController,
		WeatherController:     weatherController,
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
