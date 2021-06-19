package main

import (
	_authHttpDeliver "github.com/creamyshit/gologin/domain/auth/delivery/http"
	_authRepo "github.com/creamyshit/gologin/domain/auth/repository"
	_authUcase "github.com/creamyshit/gologin/domain/auth/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	dsn := "host=localhost user=postgres password=Postgres123!@# dbname=pos port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	ar := _authRepo.AuthRepository(db)
	au := _authUcase.AuthUsecase(ar)

	_authHttpDeliver.AuthHandler(app, au)

	app.Listen(":3000")
}
