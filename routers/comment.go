package routers

import (
	CommentAdapter "github.com/Farenthigh/Fitbuddy-BE/adapters/comment"
	CommentUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/comment"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitCommentRouter(app *fiber.App, db *gorm.DB) {
	commentRepo := CommentAdapter.NewCommentGorm(db)
	commentUsecase := CommentUsecase.NewCommentService(commentRepo)
	commentHandler := CommentAdapter.NewCommentAdapter(commentUsecase)

	comment := app.Group("/comments")
	comment.Get("/", commentHandler.GetAll)
	comment.Get("/:id", commentHandler.GetByID)
	comment.Post("/", commentHandler.Create)
	comment.Put("/:id", commentHandler.Update)
	comment.Delete("/:id", commentHandler.DeleteByID)
}
