package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"

	// "github.com/gofiber/fiber/v2/middleware/jwt"

	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/model"
)

func AuthenticateJWT(config configuration.Config) func(*fiber.Ctx) error {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			userID := int(claims["id"].(float64))
			email := claims["email"].(string)
			noTelp := claims["noTelp"].(string)
			ctx.Locals("id", userID)
			ctx.Locals("email", email)
			ctx.Locals("noTelp", noTelp)

			return ctx.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
					Status:  false,
					Message: "Bad Request",
					Errors:  "Invalid token",
					Data:    "Missing or malformed JWT",
				})
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
					Status:  false,
					Message: "Unauthorized",
					Errors:  "Invalid token",
					Data:    "Invalid or expired JWT",
				})
			}
		},
	})
}
