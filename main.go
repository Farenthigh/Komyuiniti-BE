package main

import (
	"fmt"

	"github.com/Farenthigh/Fitbuddy-BE/config"
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	"github.com/Farenthigh/Fitbuddy-BE/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".env")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbSchema)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3001", // หรือ "*" ชั่วคราว (ไม่แนะนำใน production)
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))
	Entities.Init(db)
	fmt.Println(config.Port)

	routers.InitUserRoute(app, db)
	routers.InitTweetRoute(app, db)
	routers.InitEventRoute(app, db)
	routers.InitAnimeRoute(app, db)
	routers.InitReviewRoute(app, db)
	routers.InitCommentRouter(app, db)
	routers.InitFavoriteRouter(app, db)

	app.Listen(":80")
}
