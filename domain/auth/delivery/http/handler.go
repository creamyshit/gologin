package http

import (
	"github.com/creamyshit/gologin/domain/auth"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AuthUsecase auth.Usecase
}

func AuthHandler(app *fiber.App, us auth.Usecase) {

	handler := &Handler{
		AuthUsecase: us,
	}

	app.Post("/signup", handler.SignUp)
	app.Post("/login", handler.Signin)
}

func (a *Handler) Signin(c *fiber.Ctx) error {
	payload := &auth.SignUpPayload{}
	err := c.BodyParser(payload)

	if err != nil {
		return err
	}

	res, err := a.AuthUsecase.Signin(payload)

	if err != nil {
		return c.JSON(&fiber.Map{
			"success": true,
			"return":  err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"return":  res,
	})
}

func (a *Handler) SignUp(c *fiber.Ctx) error {
	payload := &auth.SignUpPayload{}
	err := c.BodyParser(payload)

	if err != nil {
		return err
	}

	res, err := a.AuthUsecase.Signup(payload)

	return c.JSON(&fiber.Map{
		"success": true,
		"return":  res,
	})
}
