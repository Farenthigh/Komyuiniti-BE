package FavoriteAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	FavoriteUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/favorite"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
)

type FavoriteHandler struct {
	favoriteUsecase FavoriteUsecase.FavoriteUsecase
}

func NewFavoriteAdapter(favoriteUsecase FavoriteUsecase.FavoriteUsecase) *FavoriteHandler {
	return &FavoriteHandler{
		favoriteUsecase: favoriteUsecase,
	}
}
func (h *FavoriteHandler) GetAll(c *fiber.Ctx) error {
	favorites, err := h.favoriteUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", favorites)
}
func (h *FavoriteHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	favoriteID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	favorite, err := h.favoriteUsecase.GetByID(&favoriteID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", favorite)
}
func (h *FavoriteHandler) Create(c *fiber.Ctx) error {
	var favorite Entities.Favorite
	if err := c.BodyParser(&favorite); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	if err := h.favoriteUsecase.Create(&favorite); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Success", "", favorite)
}
func (h *FavoriteHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	favoriteID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}

	var favorite Entities.Favorite
	if err := c.BodyParser(&favorite); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	favorite.ID = favoriteID

	if err := h.favoriteUsecase.Update(&favorite); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", favorite)
}
func (h *FavoriteHandler) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	favoriteID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}

	if err := h.favoriteUsecase.DeleteByID(&favoriteID); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", nil)
}
func (h *FavoriteHandler) GetByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	if userID == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "User ID is required", "User ID is required", nil)
	}
	userIDUint, err := utils.StringToUint(userID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid User ID format", err.Error(), nil)
	}
	favorites, err := h.favoriteUsecase.GetByUserID(&userIDUint)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", favorites)
}