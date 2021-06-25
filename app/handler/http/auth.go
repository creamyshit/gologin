package handler

import (
	"github.com/creamyshit/gologin/domain"
	"github.com/creamyshit/gologin/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AuthUsecase domain.AuthEntity
}

func AuthHandler(app *fiber.App, us domain.AuthEntity) {

	handler := &Handler{
		AuthUsecase: us,
	}

	app.Post("/signup", handler.SignUp)
	app.Post("/login", handler.Signin)
}

func (a *Handler) Signin(c *fiber.Ctx) error {

	payload := &domain.SignUpPayload{}
	if err := c.BodyParser(payload); err != nil {
		return utils.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	validation := validator.New()
	if err := validation.Struct(payload); err != nil {
		return utils.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	res, err := a.AuthUsecase.Signin(payload)
	if err != nil {
		return utils.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	return utils.AuthResponse(c, true, res, "your custom msg here", 200, "")
}

func (a *Handler) SignUp(c *fiber.Ctx) error {

	payload := &domain.SignUpPayload{}
	if err := c.BodyParser(payload); err != nil {
		return utils.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	validation := validator.New()
	if err := validation.Struct(payload); err != nil {
		return utils.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	res, err := a.AuthUsecase.Signup(payload)
	if err != nil {
		return utils.AuthResponse(c, false, err.Error(), "your custom msg here", 400, "")
	}

	return utils.AuthResponse(c, true, res, "your custom msg here", 200, "")
}
