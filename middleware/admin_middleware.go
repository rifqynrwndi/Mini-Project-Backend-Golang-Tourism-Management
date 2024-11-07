package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil nilai "role" dari token JWT
		role := c.Get("userRole")

		// Periksa apakah peran adalah "admin"
		if role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "Access forbidden: admin only",
			})
		}

		return next(c)
	}
}
