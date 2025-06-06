package AnimeAdapter

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/Farenthigh/Fitbuddy-BE/config"
	Entities "github.com/Farenthigh/Fitbuddy-BE/entities"
	AnimeUsecase "github.com/Farenthigh/Fitbuddy-BE/usecases/anime"
	"github.com/Farenthigh/Fitbuddy-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/api/option"
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

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "failed to parse form", err.Error(), nil)
	}
	if form == nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "failed to parse form", "failed to parse form", nil)
	}

	if form.Value["title"] != nil {
		anime.Title = form.Value["title"][0]
	}
	if form.Value["description"] != nil {
		anime.Description = form.Value["description"][0]
	}
	if form.Value["genre"]!= nil {
		anime.Genres = form.Value["genre"][0]
	}
	if form.Value["Status"] != nil {
		anime.Status = form.Value["Status"][0]
	}
	if form.Value["studio"]!= nil {
		anime.Studio = form.Value["studio"][0]
	}
	if form.Value["episodes"] != nil {
		ep := form.Value["episodes"][0]
		episodes, err := strconv.Atoi(ep)
		if err != nil {
			return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid episodes format", err.Error(), nil)
		}
		anime.Episodes = episodes
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.BucketKey))
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create storage client", err.Error(), nil)
	}
	defer client.Close()

	files := form.File["anime_image"]
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
		anime.Anime_image = shareableURLs[0]
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
