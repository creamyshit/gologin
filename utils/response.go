package utils

import (
	"time"

	"github.com/creamyshit/gologin/src/model"

	"github.com/gofiber/fiber/v2"
)

type authResponse struct {
	Result       interface{} `json:"result"`
	Message      string      `json:"message"`
	Success      bool        `json:"success"`
	SessionToken string      `json:"sessiontoken"`
}

func AuthResponse(c *fiber.Ctx, success bool, data interface{}, message string, code int, token string) error {
	return c.Status(code).JSON(&authResponse{
		Success:      success,
		Result:       data,
		Message:      message,
		SessionToken: token,
	})
}

func subtractTime(time1, time2 time.Time) float64 {
	diff := time2.Sub(time1).Seconds()
	return diff
}

func HideCredential(a *model.Auth) *model.Auth {
	a.Password = ""
	a.Salt = nil
	return a
}
