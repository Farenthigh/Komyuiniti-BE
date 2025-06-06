package ReviewAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	ReviewUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/review"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	reviewUsecase ReviewUsecase.ReviewUsecase
}

func NewReviewAdapter(reviewUsecase ReviewUsecase.ReviewUsecase) *ReviewHandler {
	return &ReviewHandler{
		reviewUsecase: reviewUsecase,
	}
}
func (h *ReviewHandler) GetAll(c *fiber.Ctx) error {
	reviews, err := h.reviewUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", reviews)
}
func (h *ReviewHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	reviewID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	review, err := h.reviewUsecase.GetByID(&reviewID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", review)
}

func (h *ReviewHandler) Create(c *fiber.Ctx) error {
	var review Entities.Review
	if err := c.BodyParser(&review); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	if err := h.reviewUsecase.Create(&review); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Success", "", review)
}
func (h *ReviewHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	reviewID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	var review Entities.Review
	if err := c.BodyParser(&review); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	review.ID = reviewID
	if err := h.reviewUsecase.Update(&review); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", review)
}
func (h *ReviewHandler) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	reviewID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	if err := h.reviewUsecase.DeleteByID(&reviewID); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", nil)
}

func (h *ReviewHandler) GetByAnimeID(c *fiber.Ctx) error {
	id := c.Params("animeID")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Anime ID is required", "Anime ID is required", nil)
	}
	animeID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid Anime ID format", err.Error(), nil)
	}
	reviews, err := h.reviewUsecase.GetByAnimeID(&animeID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", reviews)
}