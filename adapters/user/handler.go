package UserAdapter

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"cloud.google.com/go/storage"
	"github.com/Farenthigh/Fitbuddy-BE/config"
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	UserModels "github.com/Farenthigh/Fitbuddy-BE/model/user"
	UserUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/user"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/api/option"
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
func (a *UserHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var user Entities.User

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "failed to parse form", err.Error(), nil)
	}
	if form == nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "failed to parse form", "failed to parse form", nil)
	}
	if form.Value["username"] != nil {
		user.UserName = form.Value["username"][0]
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.BucketKey))
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create storage client", err.Error(), nil)
	}
	defer client.Close()

	files := form.File["user_image"]
	if len(files) >= 1 {
		bucketName := config.BucketName
		shareableURLs := make([]string, 0, len(files))

		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Error opening file", err.Error(), nil)
			}

			objectName := fmt.Sprintf("images/%d_%s", time.Now().UnixNano(), fileHeader.Filename)
			token := uuid.New().String()

			wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
			wc.Metadata = map[string]string{
				"firebaseStorageDownloadTokens": token,
			}

			if _, err = io.Copy(wc, file); err != nil {
				file.Close()
				wc.Close()
				return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to write file to bucket", err.Error(), nil)
			}
			file.Close()
			if err := wc.Close(); err != nil {
				return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to close writer", err.Error(), nil)
			}

			shareableURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s",
				bucketName, url.QueryEscape(objectName), token)
			shareableURLs = append(shareableURLs, shareableURL)
		}
		user.UserImage = shareableURLs[0]
	}
	fmt.Println("user.UserImage", user.UserImage)

	if userID == nil {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", "Unauthorized", nil)
	}
	userIDFloat, ok := userID.(float64)
	if !ok {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Invalid user ID type", "Invalid user ID type", nil)
	}
	userIDUint := uint(userIDFloat)
	fmt.Println("form", user.UserName)
	msg, err := a.UserUsecase.Update(&userIDUint, &user)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "user update failed", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "User updated", "", msg)
}

func (a *UserHandler) GetMyTweets(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", "Unauthorized", nil)
	}
	userIDFloat, ok := userID.(float64)
	if !ok {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Invalid user ID type", "Invalid user ID type", nil)
	}
	userIDUint := uint(userIDFloat)

	tweets, err := a.UserUsecase.GetMyTweets(&userIDUint)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Failed to get tweets", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", tweets)
}

func (a *UserHandler) GetMyEvents(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", "Unauthorized", nil)
	}
	userIDFloat, ok := userID.(float64)
	if !ok {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Invalid user ID type", "Invalid user ID type", nil)
	}
	userIDUint := uint(userIDFloat)

	events, err := a.UserUsecase.GetMyEvents(&userIDUint)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Failed to get events", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", events)
}