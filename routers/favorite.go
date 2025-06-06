package routers

import (
	FavoriteAdapter "github.com/Farenthigh/Fitbuddy-BE/adapters/favorite"
	FavoriteUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/favorite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitFavoriteRouter(app *fiber.App, db *gorm.DB) {
	favoriteRepo := FavoriteAdapter.NewFavoriteGorm(db)
	favoriteUsecase := FavoriteUsecase.NewFavoriteService(favoriteRepo)
	favoriteHandler := FavoriteAdapter.NewFavoriteAdapter(favoriteUsecase)

	favorite := app.Group("/favorites")
	favorite.Get("/", favoriteHandler.GetAll)
	favorite.Get("/:id", favoriteHandler.GetByID)
	favorite.Post("/", favoriteHandler.Create)
	favorite.Put("/:id", favoriteHandler.Update)
	favorite.Delete("/:id", favoriteHandler.DeleteByID)
}