package infrafiber

import (
	"log"
	"nbid-online-shop/infra/response"
	"nbid-online-shop/internal/config"
	"nbid-online-shop/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("token invalid")
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		token := bearer[1]

		publicId, role, err := utility.ValidateToken(token, config.Cfg.App.Encryption.JWTSecret)
		if err != nil {
			log.Println(err.Error())
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		c.Locals("ROLE", role)
		c.Locals("PUBLIC_ID", publicId)

		return c.Next()
	}
}
