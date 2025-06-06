package tweetAdapter

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"cloud.google.com/go/storage"
	"github.com/Farenthigh/Fitbuddy-BE/config"
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	tweetUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/tweet"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type TweetHandler struct {
	TweetUsecase tweetUsecase.TweetUsecase
}

func NewTweetAdapter(tweetUsecase tweetUsecase.TweetUsecase) TweetHandler {
	return TweetHandler{
		TweetUsecase: tweetUsecase,
	}
}
func (a *TweetHandler) GetAll(c *fiber.Ctx) error {
	tweets, err := a.TweetUsecase.GetAll()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", tweets)
}
func (a *TweetHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	tweetID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	tweet, err := a.TweetUsecase.GetByID(&tweetID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", tweet)
}
func (a *TweetHandler) Create(c *fiber.Ctx) error {
	var tweet Entities.Tweet

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Failed to parse form", err.Error(), nil)
	}
	if form == nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Form data is missing", "Form data is missing", nil)
	}
	if form.Value["description"] != nil {
		tweet.Description = form.Value["description"][0]
	}
	if form.Value["author_id"] == nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Author ID is required", "Author ID is required", nil)
	}
	authorID, err := utils.StringToUint(form.Value["author_id"][0])
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid Author ID format", err.Error(), nil)
	}
	tweet.AuthorID = authorID

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.BucketKey))
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create storage client", err.Error(), nil)
	}
	defer client.Close()

	files := form.File["tweet_image"]
	if len(files) > 0 {
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
		tweet.Tweet_image = shareableURLs[0]
	}
	if err := a.TweetUsecase.Create(&tweet); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Tweet created successfully", "", tweet)
}

func (a *TweetHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	tweetID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	var tweet Entities.Tweet
	if err := c.BodyParser(&tweet); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	tweet.ID = tweetID
	if err := a.TweetUsecase.Update(&tweet); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", tweet)
}
func (a *TweetHandler) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID is required", "ID is required", nil)
	}
	tweetID, err := utils.StringToUint(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	if err := a.TweetUsecase.DeleteByID(&tweetID); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, err.Error(), err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Success", "", nil)
}
