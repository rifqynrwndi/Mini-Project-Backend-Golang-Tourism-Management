package routes

import (
	"os"
	"tourism-monitoring/controllers/ai"
	"tourism-monitoring/controllers/auth"
	"tourism-monitoring/controllers/places"
	"tourism-monitoring/controllers/tourists"
	"tourism-monitoring/controllers/trash_report"
	"tourism-monitoring/controllers/visit_report"
	"tourism-monitoring/controllers/weather"
	"tourism-monitoring/middleware"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController        *auth.AuthController
	TouristsController    *tourists.TouristsController
	PlacesController      *places.PlacesController
	VisitReportController *visit_report.VisitReportController
	TrashReportController *trash_report.TrashReportController
	AIController          *ai.AIController
	WeatherController     *weather.WeatherController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	// Authentication routes
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/register", rc.AuthController.RegisterController)

	// Protected routes with JWT
	eJWT := e.Group("")
	eJWT.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey: "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middleware.JwtCustomClaims)
		},
	}))

	// Admin-only routes for tourists
	eAdmin := eJWT.Group("/tourists")
	eAdmin.Use(middleware.AdminOnly)
	eAdmin.GET("", rc.TouristsController.GetAllTourists)
	eAdmin.GET("/:id", rc.TouristsController.GetTouristByID)
	eAdmin.POST("", rc.TouristsController.InsertTourist)
	eAdmin.PUT("/:id", rc.TouristsController.UpdateTourist)
	eAdmin.DELETE("/:id", rc.TouristsController.DeleteTourist)

	// Public routes for tourists
	e.GET("/tourists", rc.TouristsController.GetAllTourists)
	e.GET("/tourists/:id", rc.TouristsController.GetTouristByID)

	// Admin-only routes for places
	eAdminPlaces := eJWT.Group("/places")
	eAdminPlaces.Use(middleware.AdminOnly)
	eAdminPlaces.GET("", rc.PlacesController.GetAllPlaces)
	eAdminPlaces.GET("/:id", rc.PlacesController.GetPlaceById)
	eAdminPlaces.POST("", rc.PlacesController.InsertPlace)
	eAdminPlaces.PUT("/:id", rc.PlacesController.UpdatePlace)
	eAdminPlaces.DELETE("/:id", rc.PlacesController.DeletePlace)

	// Public routes for places
	e.GET("/places", rc.PlacesController.GetAllPlaces)
	e.GET("/places/:id", rc.PlacesController.GetPlaceById)

	// Admid-only routes for visit report
	eAdminVisitReport := eJWT.Group("/visit_reports")
	eAdminVisitReport.Use(middleware.AdminOnly)
	eAdminVisitReport.GET("", rc.VisitReportController.GetAllVisitReports)
	eAdminVisitReport.GET("/:id", rc.VisitReportController.GetVisitReportByID)
	eAdminVisitReport.POST("", rc.VisitReportController.InsertVisitReport)
	eAdminVisitReport.PUT("/:id", rc.VisitReportController.UpdateVisitReport)
	eAdminVisitReport.DELETE("/:id", rc.VisitReportController.DeleteVisitReport)

	// Admin-only routes for trash report
	eAdminTrashReport := eJWT.Group("/trash_reports")
	eAdminTrashReport.Use(middleware.AdminOnly)
	eAdminTrashReport.GET("/places/:id", rc.TrashReportController.GetTrashReportByPlaceID)
	eAdminTrashReport.GET("/:id", rc.TrashReportController.GetTrashReportByID)
	eAdminTrashReport.POST("", rc.TrashReportController.InsertTrashReport)
	eAdminTrashReport.PUT("/:id", rc.TrashReportController.UpdateTrashReport)
	eAdminTrashReport.DELETE("/:id", rc.TrashReportController.DeleteTrashReport)

	// Public routes for AI Prediction
	e.GET("/predict", rc.AIController.PredictVisitsAndRecommend)

	// Public routes for weather
	e.GET("/weather", rc.WeatherController.GetWeather)
}
