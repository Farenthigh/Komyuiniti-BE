package main

import (
	"fmt"

	"github.com/Farenthigh/Fitbuddy-BE/config"
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	"github.com/Farenthigh/Fitbuddy-BE/routers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbSchema)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	app := fiber.New()
	Entities.Init(db)

	routers.InitUserRoute(app, db)
	routers.InitTweetRoute(app, db)

	app.Listen(fmt.Sprintf(":%s", config.Port))
}
