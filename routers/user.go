package routers

import (
	UserAdapter "github.com/Farenthigh/Fitbuddy-BE/adapters/user"
	UserUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/user"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitUserRoute(app *fiber.App, db *gorm.DB) {
	userRepo := UserAdapter.NewUserGorm(db)
	userUsecase := UserUsecase.NewUserService(userRepo)
	userHandler := UserAdapter.NewUserAdapter(userUsecase)

	user := app.Group("/users")
	user.Post("/", userHandler.Insert)
	user.Get("/", userHandler.GetAll)
	user.Post("/login", userHandler.Login)

	protected := user.Group("/")
	protected.Use(utils.IsExist)
	protected.Get("/me", userHandler.Me)
	protected.Put("/me", userHandler.Update)
	protected.Get("/tweets", userHandler.GetMyTweets)
	protected.Get("/events", userHandler.GetMyEvents)

	protected.Get("/logout", userHandler.Logout)
}
