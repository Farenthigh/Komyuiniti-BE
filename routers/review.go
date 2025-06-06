package routers

import (
	ReviewAdapter "github.com/Farenthigh/Fitbuddy-BE/adapters/review"
	ReviewUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/review"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitReviewRoute(app *fiber.App, db *gorm.DB) {
	reviewRepo := ReviewAdapter.NewReviewGorm(db)
	reviewUsecase := ReviewUsecase.NewReviewService(reviewRepo)
	ReviewHandler := ReviewAdapter.NewReviewAdapter(reviewUsecase)

	review := app.Group("/reviews")
	review.Get("/", ReviewHandler.GetAll)
	review.Get("/:id", ReviewHandler.GetByID)
	review.Post("/", ReviewHandler.Create)
	review.Put("/:id", ReviewHandler.Update)
	review.Delete("/:id", ReviewHandler.DeleteByID)
	review.Get("/anime/:animeID", ReviewHandler.GetByAnimeID)
}
