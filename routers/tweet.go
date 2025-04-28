package routers

import (
	tweetAdapter "github.com/Farenthigh/Fitbuddy-BE/adapters/tweet"
	tweetUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/tweet"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitTweetRoute(app *fiber.App, db *gorm.DB) {
	tweetRepo := tweetAdapter.NewTweetGorm(db)
	tweetUsecase := tweetUsecase.NewTweetService(tweetRepo)
	tweetHandler := tweetAdapter.NewTweetAdapter(tweetUsecase)

	tweet := app.Group("/tweets")
	tweet.Get("/", tweetHandler.GetAll)
	tweet.Get("/:id", tweetHandler.GetByID)
	tweet.Post("/", tweetHandler.Create)
	tweet.Put("/:id", tweetHandler.Update)
	tweet.Delete("/:id", tweetHandler.DeleteByID)

}
