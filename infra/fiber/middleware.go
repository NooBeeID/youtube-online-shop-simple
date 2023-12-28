package infrafiber

import (
	"context"
	"fmt"
	"log"
	"nbid-online-shop/infra/response"
	"nbid-online-shop/internal/config"
	infraLog "nbid-online-shop/internal/log"
	"nbid-online-shop/utility"
	"strings"
	"time"

	"github.com/NooBeeID/go-logging/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Trace() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		now := time.Now()
		traceId := uuid.New()
		c.Set("X-Trace-ID", traceId.String())

		data := map[logger.LogKey]interface{}{
			logger.TRACER_ID: traceId,
			logger.METHOD:    c.Route().Method,
			logger.PATH:      string(c.Context().URI().Path()),
		}

		ctx = context.WithValue(ctx, logger.DATA, data)

		// get request
		infraLog.Log.Infof(ctx, "incoming request")

		c.SetUserContext(ctx)
		err := c.Next()

		// finish request
		data[logger.RESPONSE_TIME] = time.Since(now).Milliseconds()
		data[logger.RESPONSE_TYPE] = "ms"

		httpStatusCode := c.Response().Header.StatusCode()
		if httpStatusCode >= 200 && httpStatusCode <= 299 {
			ctx = context.WithValue(ctx, logger.DATA, data)
			infraLog.Log.Infof(ctx, "success")
		} else {
			respBody := c.Response().Body()
			data["response_body"] = fmt.Sprintf("%s", respBody)

			ctx = context.WithValue(ctx, logger.DATA, data)
			infraLog.Log.Errorf(ctx, "success")

		}

		return err
	}
}

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

func CheckRoles(authorizedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := fmt.Sprintf("%v", c.Locals("ROLE"))

		isExists := false
		for _, authorizedRole := range authorizedRoles {
			if role == authorizedRole {
				isExists = true
				break
			}
		}

		if !isExists {
			return NewResponse(
				WithError(response.ErrorForbiddenAccess),
			).Send(c)
		}

		return c.Next()
	}
}
