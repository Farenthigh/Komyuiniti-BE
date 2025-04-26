package UserAdapter

import (
	UserModels "github.com/Farenthigh/Fitbuddy-BE/model/user"
	UserUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/user"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUsecase UserUsecase.UserUsecase
}

func NewUserAdapter(userUsecase UserUsecase.UserUsecase) UserHandler {
	return UserHandler{
		UserUsecase: userUsecase,
	}
}

func (a *UserHandler) Insert(c *fiber.Ctx) error {
	var user UserModels.RegisterInput
	if err := c.BodyParser(&user); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "failed to parse body", err.Error(), nil)
	}
	msg, err := a.UserUsecase.Register(&user)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, msg, err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "User created", "", user)
}

func (a *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := a.UserUsecase.GetAll()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", users)
}

func (a *UserHandler) Login(c *fiber.Ctx) error {
	var user UserModels.LoginInput
	if err := c.BodyParser(&user); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "failed to parse body", err.Error(), nil)
	}
	msg, err := a.UserUsecase.Login(&user)
	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    msg,
		MaxAge:   3600,
		Secure:   true,
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteStrictMode,
	}
	c.Cookie(&cookie)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, msg, err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Login success", "", msg)
}

func (a *UserHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", "Unauthorized", nil)
	}
	userIDFloat, ok := userID.(float64)
	if !ok {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Invalid user ID type", "Invalid user ID type", nil)
	}
	userIDUint := uint(userIDFloat)
	user, err := a.UserUsecase.Me(&userIDUint)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", user)
}

func (a *UserHandler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    "",
		MaxAge:   -1,
		Secure:   true,
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteStrictMode,
	}
	c.Cookie(&cookie)
	return utils.ResponseJSON(c, fiber.StatusOK, "Logout success", "", nil)
}
