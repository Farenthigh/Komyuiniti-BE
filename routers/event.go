package routers

import (
	eventAdapter "github.com/Farenthigh/Fitbuddy-BE/adapters/event"
	EventUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/event"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitEventRoute(app *fiber.App, db *gorm.DB) {
	eventRepo := eventAdapter.NewEventGorm(db)
	eventUsecase := EventUsecase.NewEventService(eventRepo)
	eventHandler := eventAdapter.NewEventAdapter(eventUsecase)

	event := app.Group("/events")
	event.Get("/", eventHandler.GetAll)
	event.Get("/:id", eventHandler.GetByID)
	event.Post("/", eventHandler.Create)
	event.Put("/:id", eventHandler.Update)
	event.Delete("/:id", eventHandler.DeleteByID)
	event.Get("/user/:userID", eventHandler.GetByUserID)
}
