package routes

import (
	"os"
	"tourism-monitoring/controllers/auth"
	"tourism-monitoring/controllers/tourists"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController     *auth.AuthController
	TouristsController *tourists.TouristsController
}

func (rc RouteController) InitRoute(e *echo.Echo) {

	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/register", rc.AuthController.RegisterController)

	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	eUser := eJWT.Group("/users")
	eUser.GET("", rc.AuthController.LoginController)
	eUser.POST("", rc.AuthController.RegisterController)

	e.GET("/tourists", rc.TouristsController.GetAllTourists)
}
