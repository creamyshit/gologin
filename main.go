package main

import (
	"fmt"

	_handler "github.com/creamyshit/gologin/src/handler/http"
	"github.com/creamyshit/gologin/src/model"
	_repo "github.com/creamyshit/gologin/src/repository"
	_usecase "github.com/creamyshit/gologin/src/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	dsn := "host=localhost user=postgres password=Postgres123!@# dbname=pos port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println("Start Migration ...")
	db.AutoMigrate(&model.Auth{})

	ar := _repo.NewAuthRepository(db)
	au := _usecase.NewAuthUsecase(ar)

	_handler.AuthHandler(app, au)

	app.Listen(":3000")
}
