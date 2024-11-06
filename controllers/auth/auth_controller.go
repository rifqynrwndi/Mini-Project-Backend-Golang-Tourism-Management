package auth

import (
	"tourism-monitoring/controllers/auth/request"
	"tourism-monitoring/controllers/auth/response"
	"tourism-monitoring/controllers/base"
	"tourism-monitoring/services"

	"github.com/labstack/echo/v4"
)

func NewAuthController(as services.AuthInterface) *AuthController {
	return &AuthController{
		authService: as,
	}
}

type AuthController struct {
	authService services.AuthInterface
}

func (userController AuthController) LoginController(c echo.Context) error {
	userLogin := request.LoginRequest{}
	c.Bind(&userLogin)
	user, err := userController.authService.Login(userLogin.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromEntities(user))
}

func (userController AuthController) RegisterController(c echo.Context) error {
	userRegister := request.RegisterRequest{}
	if err := c.Bind(&userRegister); err != nil {
		return base.ErrorResponse(c, err)
	}
	user, err := userController.authService.Register(userRegister.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromEntities(user))
}
