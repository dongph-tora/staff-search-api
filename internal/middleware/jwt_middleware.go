package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"staff-search-api/pkg/jwt"
	"staff-search-api/pkg/response"
)

func JWTMiddleware(jwtService *jwt.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Unauthorized(c, "Missing authorization header.")
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return response.Unauthorized(c, "Invalid authorization header format.")
		}

		claims, err := jwtService.ValidateAccessToken(parts[1])
		if err != nil {
			return response.Unauthorized(c, "Invalid or expired token.")
		}

		c.Locals("userID", claims.Subject)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}
