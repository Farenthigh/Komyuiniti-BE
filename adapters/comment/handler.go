package CommentAdapter

import (
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	CommentUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/comment"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	commentUsecase CommentUsecase.CommentUsecase
}

func NewCommentAdapter(commentUsecase CommentUsecase.CommentUsecase) *CommentHandler {
	return &CommentHandler{
		commentUsecase: commentUsecase,
	}
}
func (h *CommentHandler) GetAll(c *fiber.Ctx) error {
	comments, err := h.commentUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", comments)
}
func (h *CommentHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	commentID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	comment, err := h.commentUsecase.GetByID(&commentID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", comment)
}
func (h *CommentHandler) Create(c *fiber.Ctx) error {
	var comment Entities.Comment
	if err := c.BodyParser(&comment); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	if err := h.commentUsecase.Create(&comment); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Success", "", comment)
}
func (h *CommentHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	commentID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}

	var comment Entities.Comment
	if err := c.BodyParser(&comment); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	comment.ID = commentID

	if err := h.commentUsecase.Update(&comment); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", comment)
}
func (h *CommentHandler) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	commentID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}

	if err := h.commentUsecase.DeleteByID(&commentID); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", nil)
}
