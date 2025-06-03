package eventAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	EventUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/event"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	EventUsecase EventUsecase.EventUsecase
}

func NewEventAdapter(eventUsecase EventUsecase.EventUsecase) *EventHandler {
	return &EventHandler{
		EventUsecase: eventUsecase,
	}
}
func (h *EventHandler) GetAll(c *fiber.Ctx) error {
	events, err := h.EventUsecase.GetAll()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", events)
}
func (h *EventHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	eventID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	event, err := h.EventUsecase.GetByID(&eventID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", event)
}
func (h *EventHandler) Create(c *fiber.Ctx) error {
	var event Entities.Event
	if err := c.BodyParser(&event); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	if err := h.EventUsecase.Create(&event); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Success", "", event)
}
func (h *EventHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	eventID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	var event Entities.Event
	if err := c.BodyParser(&event); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	event.ID = eventID
	if err := h.EventUsecase.Update(&event); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", event)
}
func (h *EventHandler) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	eventID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	if err := h.EventUsecase.DeleteByID(&eventID); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", nil)
}
func (h *EventHandler) GetByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	if userID == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "User ID is required", "User ID is required", nil)
	}
	userIDUint, err := utils.StringToUint(userID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid User ID format", err.Error(), nil)
	}
	events, err := h.EventUsecase.GetByUserID(&userIDUint)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", events)
}
