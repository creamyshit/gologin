package http

import (
	"github.com/creamyshit/gologin/domain/auth"
	"github.com/creamyshit/gologin/helper"
	"github.com/go-playground/validator/v10"
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
		return helper.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	validation := validator.New()
	err = validation.Struct(payload)

	if err != nil {
		return helper.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	res, err := a.AuthUsecase.Signin(payload)

	if err != nil {
		return helper.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	return helper.AuthResponse(c, true, res, "your custom msg here", 200, "")
}

func (a *Handler) SignUp(c *fiber.Ctx) error {
	payload := &auth.SignUpPayload{}
	err := c.BodyParser(payload)

	if err != nil {
		return err
	}

	validation := validator.New()
	err = validation.Struct(payload)

	if err != nil {
		return helper.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	res, err := a.AuthUsecase.Signup(payload)

	if err != nil {
		return helper.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	return helper.AuthResponse(c, true, res, "your custom msg here", 200, "")
}
