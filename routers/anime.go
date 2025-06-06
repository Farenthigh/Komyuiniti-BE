package routers

import (
	AnimeAdapter "github.com/Farenthigh/Fitbuddy-BE/adapters/anime"
	AnimeUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/anime"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitAnimeRoute(app *fiber.App, db *gorm.DB) {
	animeRepo := AnimeAdapter.NewAnimeGorm(db)
	animeUsecase := AnimeUsecase.NewAnimeService(animeRepo)
	animeHandler := AnimeAdapter.NewAnimeAdapter(animeUsecase)

	anime := app.Group("/animes")
	anime.Get("/", animeHandler.GetAll)
	anime.Get("/:id", animeHandler.GetByID)
	anime.Post("/", animeHandler.Create)
	anime.Put("/:id", animeHandler.Update)
	anime.Delete("/:id", animeHandler.DeleteByID)
}
