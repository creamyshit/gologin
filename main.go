package main

import (
	"fmt"

	_entity "github.com/creamyshit/gologin/app/entity"
	_handler "github.com/creamyshit/gologin/app/handler/http"
	_repo "github.com/creamyshit/gologin/app/repository"
	"github.com/creamyshit/gologin/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	dsn := "host=localhost user=postgres password=Postgres123!@# dbname=pos port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println("Start Migration ...")
	db.AutoMigrate(&domain.Auth{})

	ar := _repo.NewAuthRepository(db)
	au := _entity.NewAuthEntity(ar)

	_handler.AuthHandler(app, au)

	app.Listen(":3000")
}
