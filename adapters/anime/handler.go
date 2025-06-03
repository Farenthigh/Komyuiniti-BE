package AnimeAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	AnimeUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/anime"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
)

type AnimeHandler struct {
	AnimeUsecase AnimeUsecase.AnimeUsecase
}

func NewAnimeAdapter(animeUsecase AnimeUsecase.AnimeUsecase) *AnimeHandler {
	return &AnimeHandler{
		AnimeUsecase: animeUsecase,
	}
}
func (h *AnimeHandler) GetAll(c *fiber.Ctx) error {
	animes, err := h.AnimeUsecase.GetAll()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", animes)
}
func (h *AnimeHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	animeID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	anime, err := h.AnimeUsecase.GetByID(&animeID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", anime)
}
func (h *AnimeHandler) Create(c *fiber.Ctx) error {
	var anime Entities.Anime
	if err := c.BodyParser(&anime); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	if err := h.AnimeUsecase.Create(&anime); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Success", "", anime)
}
func (h *AnimeHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	animeID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	var anime Entities.Anime
	if err := c.BodyParser(&anime); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	anime.ID = animeID
	if err := h.AnimeUsecase.Update(&anime); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", anime)
}
func (h *AnimeHandler) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	animeID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	if err := h.AnimeUsecase.DeleteByID(&animeID); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Delete Success", "", nil)
}
func (h *AnimeHandler) GetByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	if userID == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "User ID is required", "User ID is required", nil)
	}
	UserIDUint, err := utils.StringToUint(userID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid User ID format", err.Error(), nil)
	}
	animes, err := h.AnimeUsecase.GetByUserID(&UserIDUint)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", animes)
}
